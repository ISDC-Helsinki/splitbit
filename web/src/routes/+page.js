import { error } from '@sveltejs/kit';
import { PUBLIC_SERVER_URL } from '$env/static/public'

export async function load({ }) {
        let fetchJson = await fetch(PUBLIC_SERVER_URL + "/groups-nonauthed");
        if (fetchJson.ok) {
                let data = await fetchJson.json();
                return { groups: data }
        } else {
                error(404, 'Not found');
        }

}
