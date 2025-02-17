import api from "$lib/api/api";
import { error as errorOut } from "@sveltejs/kit";

export async function load({ params }) {
  const {
    data, // only present if 2XX response
    error,
    response,
  } = await api.GET("/groups/{id}", { params: { path: { id: params.slug } } });
  if (response.status != 200) {
    errorOut(response.status, response.statusText);
  }
  return {
    items: data!.items,
    group_id: params.slug,
    balance: data?.money_balance,
  };
}
