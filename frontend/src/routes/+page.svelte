<script lang="ts">
	import type { ActionData } from './$types';

	let { form }: { form: ActionData } = $props();
</script>

<h1>Counter</h1>

<p>Value: {form?.value ?? 'N/A'}</p>

<form method="POST">
	<button type="submit">Increment</button>
</form>
