import { error as errorOut } from '@sveltejs/kit';
import api from '$lib/api/api';


export async function load({ }) {
        const {
                data, // only present if 2XX response
                error, // only present if 4XX or 5XX response
                response
        } = await api.GET("/groups-nonauthed")
        if (response.status != 200) {
                errorOut(response.status, response.statusText);
        }
        return { groups: data }

}
