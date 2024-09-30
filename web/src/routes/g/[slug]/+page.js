import { error } from '@sveltejs/kit';
import { PUBLIC_SERVER_URL } from '$env/static/public'

export async function load({ params }) {
        let group_info = await fetch(`${PUBLIC_SERVER_URL}/groups/${params.slug}`);
        let items = await fetch(`${PUBLIC_SERVER_URL}/groups/${params.slug}/items`);
        if (items.ok && group_info.ok) {
                let item_data = await items.json();
                let group_data = await group_info.json();
                return { group_id: params.slug, items: item_data, group_data: group_data }
        } else {
                error(404, 'Not found');
        }

}
