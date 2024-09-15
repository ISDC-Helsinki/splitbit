import { error } from '@sveltejs/kit';
import { PUBLIC_SERVER_URL } from '$env/static/public'

export async function load({ params }) {
        let fetchJson = await fetch(`${PUBLIC_SERVER_URL}/groups/${params.slug}/items`);
        if (fetchJson.ok) {
                let data = await fetchJson.json();
                return { group_id: params.slug, items: data }
        } else {
                error(404, 'Not found');
        }

}
