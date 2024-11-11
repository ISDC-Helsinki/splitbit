<script lang="ts">
  import Button from "$lib/buttons/Button.svelte";
  import TextFieldOutlined from "$lib/forms/TextFieldOutlined.svelte";
  import Icon from "$lib/misc/_icon.svelte";
  import add from "@ktibow/iconset-material-symbols/add";
  import close from "@ktibow/iconset-material-symbols/close";

  let name: string = "Svelte";
  let src: string = "/tutorial/image.gif";
  let count: number = 0;
  let newMember: string = "";
  let members: string[] = [];

  function handleClick() {
    count += 1;
  }

  function addMember() {
    if (newMember.trim()) {
      members = [...members, newMember];
      newMember = "";
    }
  }

  function deleteMember(index: number) {
    members = members.filter((_, i) => i !== index);
  }
</script>

<div class="container">
  <h1 style="text-align: center;">Create Group</h1>
  <TextFieldOutlined name="Group name" />
  <h3 style="text-align: left;">Members</h3>

  <TextFieldOutlined
    trailingIcon={add}
    name="Add members"
    bind:value={newMember}
    on:trailingClick={addMember}
  />

  <ul>
    {#each members as member, index}
      <li>
        {member}
        <button type="button" on:click={() => deleteMember(index)} aria-label="Delete member">
          <Icon icon={close} />
        </button>
      </li>
    {/each}
  </ul>

    <!--Padding for the create group button-->
    <div class="spacer">

    </div>
    <Button type="filled" >Create group</Button>
  
</div>

<style>
    .spacer {
        height:1rem;
    }
    .container {
        display:flex;
        flex-direction: column;
        max-width: 40rem;
        margin: auto;
        
    }
  p {
    color: goldenrod;
    font-family: "Comic Sans MS", cursive;
    font-size: 2em;
  }
  button {
    background-color: #007bff;
    color: white;
    border: none;
    padding: 10px 20px;
    font-size: 16px;
    cursor: pointer;
    border-radius: 5px;
  }
  button:hover {
    background-color: #0056b3;
  }
  ul {
    list-style-type: none;
    padding: 0;
  }
  li {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 10px;
  }
  li button {
    background: none;
    border: none;
    padding: 0;
    cursor: pointer;
  }
</style>
