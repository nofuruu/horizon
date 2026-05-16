<script lang="ts">
	import * as Card from "$lib/components/ui/card";
	import { Input } from "$lib/components/ui/input";
	import { Button } from "$lib/components/ui/button";
	import { Label } from "$lib/components/ui/label";

	// Gunakan logika handleUpload yang kita bahas sebelumnya di sini
</script>

<div class="flex flex-col gap-4 p-4">
	<Card.Root>
		<Card.Header>
			<Card.Title>Upload Replay Match</Card.Title>
			<Card.Description
				>Unggah video pertandingan untuk dianalisis oleh hitboxMetrics.</Card.Description
			>
		</Card.Header>
		<Card.Content>
			<div class="grid w-full max-w-sm items-center gap-1.5">
				<Label for="video">Video Game (.mp4)</Label>
				<Input id="video" type="file" accept="video/*" />
			</div>
		</Card.Content>
		<Card.Footer>
			<Button>Mulai Analisis</Button>
		</Card.Footer>
	</Card.Root>
</div>
