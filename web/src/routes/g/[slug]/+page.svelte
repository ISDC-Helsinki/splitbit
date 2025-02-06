<script lang="ts">
  import { BottomSheet, Button, FAB, Icon, ListItemButton, TextField } from "$lib";
  import Fab from "$lib/buttons/FAB.svelte";
  import SideSheet from "$lib/containers/SideSheet.svelte";
  import groceryIcon from "@ktibow/iconset-material-symbols/grocery";
  import shoppingBagIcon from "@ktibow/iconset-material-symbols/shopping-bag";
  import priceIcon from "@ktibow/iconset-material-symbols/price-check";
  import check from "@ktibow/iconset-material-symbols/check";
  import { PUBLIC_SERVER_URL } from "$env/static/public";
  import ExpenseCard from "../../ExpenseCard.svelte";
  export let elaborate = true
  export let data;
  let name = "";
  let price = "";

  let addItem = (e: Event) => {
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
      data.items = [item, ...data.items];
    });
  };
  let open = true;

  let expenses = [
    {
      name: "Ppt evening",
      groupicon: "FamilyRestroom",
      expenseicon: "ShoppingCart",
      member: "Mariusz",
      expense: "Fish",
      date: 120,
    },
    {
      name: "Eetu's house",
      groupicon: "FamilyRestroom",
      expenseicon: "ShoppingCart",
      member: "Eetu",
      expense: "Decorations",
      price: 100,
    },
    {
      name: "Farm",
      groupicon: "FamilyRestroom",
      expenseicon: "Cottage",
      member: "Piettari",
      expense: "Cottage",
      price: 80,
    },
    {
      name: "Isis",
      groupicon: "FamilyRestroom",
      expenseicon: "ShoppingCart",
      member: "Jakub",
      expense: "BK entry",
      price: 1,
    },
  ];
</script>

<div>
  <h2>
    Balance:
    {data.balance}
  </h2>
</div>
<div class="screen">
  <div class="container">
    <div class="content">
      <ul>
        {#each expenses as i}
          <ExpenseCard name={i.name} expense={i.price} date={i.date} />
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
    gap: 1rem;
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
