import { error } from "@sveltejs/kit";
import { PUBLIC_SERVER_URL } from "$env/static/public";

export async function load({ params }) {
  let group_info = await fetch(`${PUBLIC_SERVER_URL}/groups/${params.slug}`);
  let items_req = await fetch(`${PUBLIC_SERVER_URL}/groups/${params.slug}/items`);

  if (group_info.ok && items_req.ok) {
    let items_data = await items_req.json();
    let group_data = await group_info.json();

    return { group_id: params.slug, items: items_data, group_data: group_data };
  } else {
    error(404, "Not found");
  }
}
