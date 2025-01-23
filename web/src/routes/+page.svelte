<script>
  import { Button } from "$lib";
  import GroupCard from "./GroupCard.svelte";
  import Icon from "$lib/misc/_icon.svelte";
  import Add from "@ktibow/iconset-material-symbols/add";
  import LongGroupCard from "./LongGroupCard.svelte";
  import FriendCard from "./FriendCard.svelte";
  import ExpenseCard from "./ExpenseCard.svelte";
  import OwedCard from "./OwedCard.svelte";
  import ArrowRightAlt from "@ktibow/iconset-material-symbols/arrow-right-alt";
  import ArrowLeftAlt from "@ktibow/iconset-material-symbols/arrow-left-alt";
  import SwapVert from "@ktibow/iconset-material-symbols/swap-vert";
  export let data;

  let owed = [
    {
      name: "Chris",
      origin: "Home Group",
      price: 120,
    },
    {
      name: "Jan",
      origin: "Eetu's trip",
      price: 100,
    },
    {
      name: "Balthazar",
      origin: "Farm",
      price: 80,
    },
  ];

  let owe = [
    {
      name: "Jakub",
      origin: "Home Group",
      price: 200,
    },
    {
      name: "Eetu",
      origin: "Eetu's trip",
      price: 300,
    },
    {
      name: "Mariusz",
      origin: "Farm",
      price: 10,
    },
  ];

  let friends = [
    {
      name: "Jakub",
      origin: "Home Group",
      price: 200,
    },
    {
      name: "Eetu",
      origin: "Eetu's trip",
      price: 300,
    },
    {
      name: "Mariusz",
      origin: "Farm",
      price: 10,
    },
  ];

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

<div class="main">
  <div class="panel">
    <div class="column">
      <div>
        <h1>Overview</h1>
      </div>

      <div>
        <div class="frame">
          <div class="small-row">
            <div class="big-icon">
              <Icon icon={SwapVert} />
            </div>

            <div>
              <div>Total balance:</div>

              <div>+60$</div>
            </div>
          </div>
        </div>

        <div class="frame">
          <div class="small-row">
            <div>
              <Icon icon={ArrowLeftAlt} />
            </div>

            <div>But you owe in total to:</div>
          </div>

          <div>
            {#each owe as g}
              <OwedCard name={g.name} origin={g.origin} price={g.price} />
            {/each}
          </div>
        </div>

        <div class="frame">
          <div class="small-row">
            <div>
              <Icon icon={ArrowRightAlt} />
            </div>

            <div>You owe in total</div>
          </div>

          <div>
            {#each owed as g}
              <OwedCard name={g.name} origin={g.origin} price={g.price} />
            {/each}
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="panel">
    <div class="column">
      <div class="top">
        <div>
          <h1>Latest expenses</h1>
        </div>

        <div class="buttons">
          <div>
            <Button type="outlined">View all</Button>
          </div>
        </div>
      </div>

      <div class="latest-expenses">
        {#each expenses as g}
          <ExpenseCard name={g.name} expense={g.expense} date={g.date} />
        {/each}
      </div>
    </div>

    <div class="column">
      <div class="top">
        <div>
          <h1>Your active groups are:</h1>
        </div>

        <div class="buttons">
          <div>
            <Button type="outlined">View all</Button>
          </div>

          <div>
            <Button type="filled">
              <div class="small-row">
                <div class="small-row">
                  <Icon icon={Add} class="Icon" />
                </div>

                <div>Add</div>
              </div>
            </Button>
          </div>
        </div>
      </div>

      <div class="active-card-holder">
        {#each data.groups as g}
          <GroupCard name={g.name} id={g.id} />
        {/each}
      </div>
    </div>

    <div class="row">
      <div class="friends">
        <div class="top">
          <div>
            <h1>Friends</h1>
          </div>

          <div class="buttons">
            <div>
              <Button type="tonal">View all</Button>
            </div>
          </div>
        </div>

        <div style="column;">
          {#each friends as g}
            <FriendCard name={g.name} />
          {/each}
        </div>
      </div>

      <div class="column">
        <div class="top">
          <div>
            <h1>Your archived groups are:</h1>
          </div>

          <div class="buttons">
            <div>
              <Button type="outlined">View all</Button>
            </div>
          </div>
        </div>

        <div class="archived-card-holder">
          {#each data.groups as g}
            <LongGroupCard name={g.name} id="{g.id}," date={g.date} />
          {/each}
        </div>
      </div>
    </div>
  </div>
</div>

<style>
  .main {
    display: flex;
    flex-direction: row;
    gap: 25px;
  }

  .active-card-holder {
    overflow: hidden;
    display: flex;
    flex-direction: row;
    gap: 15px;
  }

  .column {
    display: flex;
    flex-direction: column;
    border-radius: 10px;
    background-color: rgb(var(--m3-scheme-surface-container));
    padding: 15px;
  }

  .top {
    display: flex;
    justify-content: space-between;
    padding-left: 10px;
  }

  .buttons {
    display: flex;
    gap: 10px;
  }

  .archived-card-holder {
    overflow: hidden;
    display: flex;
    flex-direction: column;
    gap: 15px;
  }

  .panel {
    display: flex;
    flex-direction: column;
    gap: 30px;
  }

  .friends {
    width: 500px;
    border-radius: 10px;
    background-color: rgb(var(--m3-scheme-surface-container));
    padding: 15px;
  }

  .row {
    display: flex;
    flex-direction: row;
    gap: 40px;
  }

  .small-row {
    display: flex;
    flex-direction: row;
    gap: 10px;
    align-items: center;
  }


  .frame {
    border: 1px solid;
    border-radius: 10px;
    border-color: black;
    margin-top: 20px;
    padding: 10px;
    display: flex;
    flex-direction: column;
    gap: 10px;
  }

  .big-icon {
    font-size: 35px;
  }

  .latest-expenses {
    overflow: hidden;
    display: flex;
    flex-direction: column;
    gap: 15px;
    padding: 15px;
  }
</style>
