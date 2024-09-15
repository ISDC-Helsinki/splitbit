<script lang="ts">
  import { BottomSheet, Button, FAB, Icon, ListItemButton, TextField } from "$lib";
  import Fab from "$lib/buttons/FAB.svelte";
  import SideSheet from "$lib/containers/SideSheet.svelte";
  import groceryIcon from "@ktibow/iconset-material-symbols/grocery";
  import shoppingBagIcon from "@ktibow/iconset-material-symbols/shopping-bag";
  import priceIcon from "@ktibow/iconset-material-symbols/price-check";
  import check from "@ktibow/iconset-material-symbols/check";
    import { PUBLIC_SERVER_URL } from "$env/static/public";
  export let data;
  let name = "";
  let price = "";

  let addItem = (e : Event) => {
    e.preventDefault();
    const item = {
      id: 14,
      name: name,
      price: parseFloat(price),
      timestamp: Math.floor(new Date().getTime() / 1000),
      member_id: 1,
    };
    fetch(`${PUBLIC_SERVER_URL}/groups/${data.group_id}/items`, {
      method: "POST",
      body: JSON.stringify(item),
    }).then((_) => {
      data.items = [item, ...data.items]
      
    });
  };
  let open = true;

</script>

<div class="screen">
  <div class="container">
    <div class="content">
      <ul>
        {#each data.items as i}
          <ListItemButton headline={i.name} supporting="Paid something" lines={2}>
            <svelte:fragment slot="leading">
              <div class="date">
                Sep
                <b>09</b>
              </div>
              <Icon icon={groceryIcon} />
            </svelte:fragment>
            <svelte:fragment slot="trailing">
              <div class="paid m3-font-body-medium">
                You lent
                <b>{i.price}$</b>
              </div>
            </svelte:fragment>
          </ListItemButton>
        {/each}
      </ul>
    </div>
  </div>

  {#if open}
    <SideSheet on:close={() => (open = false)}>
      <form class="contents" on:submit={(e) => addItem(e)}>
        <h2 class="m3-font-headline-medium">Add item</h2>
        <div class="fields">
          <TextField leadingIcon={groceryIcon} bind:value={name} name="name" />
          <TextField leadingIcon={priceIcon} bind:value={price} name="price" />
        </div>
        <div class="fabpos">
          <FAB icon={check} size="large"></FAB>
        </div>
      </form>
    </SideSheet>
  {/if}
</div>

<style>
  .screen {
    display: flex;
    align-items: flex-start;
    height: 100%;
    gap:1rem;
  }
  .contents {
    padding: 1rem;
  }
  .fields {
    gap: 1rem;
    display: flex;
    flex-direction: column;
  }
  .fields > :global(.m3-container) {
    width: 100%;
  }
  .fabpos {
    position: absolute;
    bottom: 1rem;
    width: fit-content;
    right: 1rem;
  }
  .date {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .paid {
    margin-left: auto;
    display: flex;
    flex-direction: column;
    align-items: end;
  }
  .container {
    width: 70%;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }
  .container :global(ul button) {
    width: 100%;
  }

  .content {
    display: flex;
    flex-direction: column;
    position: relative;
    padding: 1rem;
    background-color: rgb(var(--m3-scheme-surface-container));
    color: rgb(var(--m3-scheme-on-surface));
  }
  .content:first-child {
    flex-grow: 1;
    border-radius: 1.5rem 1.5rem 0.5rem 0.5rem;
  }
  .content:last-child {
    border-radius: 0.5rem 0.5rem 1.5rem 1.5rem;
  }
</style>
