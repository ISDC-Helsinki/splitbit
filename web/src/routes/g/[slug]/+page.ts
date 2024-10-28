import api from '$lib/api/api';
import { error as errorOut } from '@sveltejs/kit';

export async function load({ params }) {
        const {
                data, // only present if 2XX response
                error, // only present if 4XX or 5XX response
                response
        } = await api.GET("/groups/{id}/items", {params:{path: { id: params.slug}}})
        if (response.status != 200) {
                errorOut(response.status, response.statusText);
        }

        return { group_id: params.slug, items: data }

}
