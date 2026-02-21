import type { Actions, PageServerLoad } from './$types';
import { createClient } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-node';
import { CounterService } from '$lib/gen/counter/v1/counter_pb';

const transport = createConnectTransport({
	baseUrl: 'http://localhost:8080',
	httpVersion: '2',
});

const client = createClient(CounterService, transport);

export const load: PageServerLoad = async () => {
	const res = await client.getValue({});
	return { value: Number(res.value) };
};

export const actions: Actions = {
	default: async () => {
		const res = await client.increment({});
		return { value: Number(res.value) };
	},
};
