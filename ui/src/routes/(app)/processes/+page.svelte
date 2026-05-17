<script lang="ts">
	import * as Card from "$lib/components/ui/card/index.js";
	import * as Table from "$lib/components/ui/table/index.js";
	import { Button } from "$lib/components/ui/button/index.js";
	import { Badge } from "$lib/components/ui/badge/index.js";
	import { Input } from "$lib/components/ui/input/index.js";
	import { Skeleton } from "$lib/components/ui/skeleton/index.js";
	import SearchIcon from "@lucide/svelte/icons/search";
	import XIcon from "@lucide/svelte/icons/x";
	import ActivityIcon from "@lucide/svelte/icons/activity";
	import { onMount, onDestroy } from "svelte";

	let processes = $state<any[]>([]);
	let filteredProcs = $state<any[]>([]);
	let searchTerm = $state("");
	let loading = $state(true);
	let intervalId: number | null = null;

	const API_BASE = "http://localhost:8080/api";

	async function fetchProcesses() {
		try {
			const res = await fetch(`${API_BASE}/processes`);
			const data = await res.json();
			processes = data.processes || [];
			filterProcesses();
			loading = false;
		} catch (err) {
			console.error(err);
			loading = false;
		}
	}

	async function killProc(pid: number) {
		if (!confirm(`Kill process ${pid}?`)) return;

		try {
			await fetch(`${API_BASE}/processes/${pid}/kill`, { method: "POST" });
			fetchProcesses();
		} catch (err) {
			alert("Failed to kill process");
		}
	}

	function filterProcesses() {
		if (!searchTerm.trim()) {
			filteredProcs = processes.slice(0, 100); // limit 100
			return;
		}
		filteredProcs = processes
			.filter(
				(p) =>
					p.name?.toLowerCase().includes(searchTerm.toLowerCase()) ||
					p.pid?.toString().includes(searchTerm)
			)
			.slice(0, 100);
	}

	$effect(() => {
		searchTerm;
		filterProcesses();
	});

	onMount(() => {
		fetchProcesses();
		intervalId = setInterval(fetchProcesses, 5000);
	});

	onDestroy(() => {
		if (intervalId) clearInterval(intervalId);
	});

	function formatTime(timestamp: number) {
		return new Date(timestamp * 1000).toLocaleString();
	}
</script>

<svelte:head>
	<title>Processes - Horizon</title>
</svelte:head>

<div class="space-y-6">
	<div>
		<h1 class="text-3xl font-bold tracking-tight">Processes</h1>
		<p class="text-muted-foreground">Running processes on the system</p>
	</div>

	<Card.Root>
		<Card.Header>
			<div class="flex items-center justify-between">
				<div class="flex items-center gap-2">
					<ActivityIcon class="size-5" />
					<Card.Title>Process List</Card.Title>
					<Badge>{processes.length} total</Badge>
				</div>
				<div class="flex items-center gap-2">
					<SearchIcon class="text-muted-foreground size-4" />
					<Input
						type="text"
						placeholder="Search by name or PID..."
						bind:value={searchTerm}
						class="w-64"
					/>
				</div>
			</div>
		</Card.Header>
		<Card.Content>
			{#if loading}
				<div class="space-y-2">
					{#each Array(10) as _}
						<Skeleton class="h-10 w-full" />
					{/each}
				</div>
			{:else}
				<Table.Root>
					<Table.Header>
						<Table.Row>
							<Table.Head>PID</Table.Head>
							<Table.Head>Name</Table.Head>
							<Table.Head>User</Table.Head>
							<Table.Head>CPU %</Table.Head>
							<Table.Head>Memory %</Table.Head>
							<Table.Head>Status</Table.Head>
							<Table.Head class="text-right">Action</Table.Head>
						</Table.Row>
					</Table.Header>
					<Table.Body>
						{#each filteredProcs as proc}
							<Table.Row>
								<Table.Cell class="font-mono text-sm">{proc.pid}</Table.Cell>
								<Table.Cell class="font-medium">{proc.name}</Table.Cell>
								<Table.Cell class="text-muted-foreground text-sm">{proc.username}</Table.Cell>
								<Table.Cell>{proc.cpu_percent?.toFixed(1) || "0.0"}%</Table.Cell>
								<Table.Cell>{proc.mem_percent?.toFixed(1) || "0.0"}%</Table.Cell>
								<Table.Cell>
									<Badge variant={proc.status === "running" ? "default" : "outline"}>
										{proc.status}
									</Badge>
								</Table.Cell>
								<Table.Cell class="text-right">
									<Button size="sm" variant="destructive" onclick={() => killProc(proc.pid)}>
										<XIcon class="size-4" />
									</Button>
								</Table.Cell>
							</Table.Row>
						{/each}
					</Table.Body>
				</Table.Root>

				{#if filteredProcs.length === 0}
					<p class="text-muted-foreground py-8 text-center">No processes found</p>
				{/if}
			{/if}
		</Card.Content>
	</Card.Root>
</div>
