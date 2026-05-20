<script lang="ts">
	import LayoutDashboardIcon from "@lucide/svelte/icons/layout-dashboard";
	import ActivityIcon from "@lucide/svelte/icons/activity";
	import ZapIcon from "@lucide/svelte/icons/zap";
	import PowerIcon from "@lucide/svelte/icons/power";

	import * as Sidebar from "$lib/components/ui/sidebar/index.js";
	import { Button } from "$lib/components/ui/button/index.js";

	type NavItem = {
		title: string;
		url: string;
		icon: typeof LayoutDashboardIcon;
		badge?: string;
	};

	type NavGroup = {
		label: string;
		items: NavItem[];
	};

	const navigation: NavGroup[] = [
		{
			label: "Monitoring",
			items: [
				{ title: "Dashboard", url: "/", icon: LayoutDashboardIcon },
				{ title: "Processes", url: "/processes", icon: ActivityIcon },
			],
		},
	];

</script>

<Sidebar.Root>
	<Sidebar.Header>
		<Sidebar.Menu>
			<Sidebar.MenuItem>
				<Sidebar.MenuButton size="lg">
					{#snippet child({ props })}
						<a href="/" {...props}>
							<div
								class="bg-primary text-primary-foreground flex aspect-square size-8 items-center justify-center rounded-lg"
							>
								<ZapIcon class="size-4" />
							</div>
							<div class="flex flex-col gap-0.5 leading-none">
								<span class="font-semibold">Horizon</span>
								<span class="text-xs">System Monitor</span>
							</div>
						</a>
					{/snippet}
				</Sidebar.MenuButton>
			</Sidebar.MenuItem>
		</Sidebar.Menu>
	</Sidebar.Header>

	<Sidebar.Content>
		{#each navigation as group (group.label)}
			<Sidebar.Group>
				<Sidebar.GroupLabel>{group.label}</Sidebar.GroupLabel>
				<Sidebar.GroupContent>
					<Sidebar.Menu>
						{#each group.items as item (item.title)}
							<Sidebar.MenuItem>
								<Sidebar.MenuButton>
									{#snippet child({ props })}
										<a href={item.url} {...props}>
											<item.icon class="size-4" />
											<span>{item.title}</span>
										</a>
									{/snippet}
								</Sidebar.MenuButton>
								{#if item.badge}
									<Sidebar.MenuBadge>{item.badge}</Sidebar.MenuBadge>
								{/if}
							</Sidebar.MenuItem>
						{/each}
					</Sidebar.Menu>
				</Sidebar.GroupContent>
			</Sidebar.Group>
		{/each}
	</Sidebar.Content>

	<Sidebar.Rail />
</Sidebar.Root>