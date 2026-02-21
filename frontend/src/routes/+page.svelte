<script lang="ts">
	import type { ActionData, PageData } from './$types';

	let { data, form }: { data: PageData; form: ActionData } = $props();
</script>

<h1>Counter</h1>

<p>Value: {form?.value ?? data.value}</p>

<form method="POST">
	<button type="submit">Increment</button>
</form>
