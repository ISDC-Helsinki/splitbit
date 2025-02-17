<script lang="ts">
    import { ListItemLabel } from "$lib";
  import api from "$lib/api/api";
  import Button from "$lib/buttons/Button.svelte";
    import Checkbox from "$lib/forms/Checkbox.svelte";
  import TextFieldOutlined from "$lib/forms/TextFieldOutlined.svelte";
  import Icon from "$lib/misc/_icon.svelte";
  import add from "@ktibow/iconset-material-symbols/add";
  import close from "@ktibow/iconset-material-symbols/close";

  let name: string = "";
  let src: string = "/tutorial/image.gif";
  let count: number = 0;
  let newMember: string = "";
  let members: {id: number; name: string}[] = [];
  let avatar: string | ArrayBuffer | null = null; // Ask Marius
  let fileinput: HTMLInputElement | null = null;

  let friends = [
    { id: 1, name: "Chris", selected: false, img: "/favicon.png" },
    { id: 2, name: "Daniel", selected: false , img: "/favicon.png"},
    { id: 3, name: "Julius", selected: false, img: "/favicon.png" }
  ];

  function handleClick() {
    count += 1;
  }

  async function addMember({}) {
    const {
      data, // only present if 2XX response
      error, // only present if 4XX or 5XX response
      response
    } = await api.GET("/members/{username}", { params: { path: { username: newMember }} });
    if (response.status != 200) {
      console.log(response.status, response.statusText);
      return;
    }
    if (newMember.trim()) {
      members = [...members, { id: data!.id, name: newMember }];
      newMember = ""; 
    }
    return members;
  }

  function deleteMember(index: number) {
    members = members.filter((_, i) => i !== index);
  }

  function toggleFriendSelection(friend: { id: number; name: string; selected: boolean }) {
    friend.selected = !friend.selected;
    if (friend.selected) {
      members = [...members, { id: friend.id, name: friend.name }];
    } else {
      members = members.filter(member => member.id !== friend.id);
    }
  }

  async function createGroup() {
    const {
      data, // only present if 2XX response
      error, // only present if 4XX or 5XX response
      response
    } = await api.POST('/groups', { body: { name: name, members: members.map((x) => x.id), icon_name: ""} })
    if (name.trim() && members.length > 0) {
        //groupCreated = true;
    } else {
      alert("Enter a group name and add at least one member.");
    }
  }

  const onFileSelected = (e: Event) => {
    const target = e.target as HTMLInputElement;
    let image;
    if (target && target.files) {
      image = target.files[0];
    }
    let reader = new FileReader();
    if (image) {
      reader.readAsDataURL(image);
    }
    reader.onload = e => {
      if (e.target) {
        avatar = e.target.result;
      }
    };
  }

  const removeAvatar = () => {
    avatar = null;
    if (fileinput) {
      fileinput.value = "";
    }
  }
</script>

<div class="container">

<div class="create-group-container">
  <h1 style="text-align: center;">Create Group</h1>
  <TextFieldOutlined 
  bind:value={name}
  name="Group name" />
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
        {member.name}
        <button type="button" on:click={() => deleteMember(index)} aria-label="Delete member">
          <Icon icon={close} />
        </button>
      </li>
    {/each}
  </ul>

  <!-- Image upload section -->
  <h3>Select the group icon</h3>
  <div class="button-group">
    <Button type="filled" on:click={() => { if (fileinput) fileinput.click(); }}>
      {#if avatar}
        Change cover photo
      {:else}
        Choose cover photo
      {/if}
    </Button>
    {#if avatar}
      <Button type="outlined" on:click={removeAvatar}>
        <Icon icon={close} />
        No cover photo</Button>
    {/if}
  </div>

  <div class="spacer"></div>

  <input style="display:none" type="file" accept=".jpg, .jpeg, .png" on:change={(e) => onFileSelected(e)} bind:this={fileinput} >

  {#if avatar}
    <img class="avatar" src={typeof avatar === 'string' ? avatar : ''} alt="d" />
  {/if}

  <!--Padding for the create group button-->
  <div class="spacer"></div>
  <div class="create-group-create-group-container">
    <Button type="filled" iconType="left" on:click={createGroup}>Create group</Button>
    
  </div>
</div>
<div class="friends-menu">
  <h3>Friends</h3>
  <ul>
    {#each friends as friend}
      <li>
      <ListItemLabel headline={friend.name} supporting={"Part of x groups"} lines={2}>
        <svelte:fragment slot="leading">
          <div class="box-wrapper">
            <img class = "friend-img" src={friend.img} alt="" />
          </div>
        </svelte:fragment>
        <svelte:fragment slot="trailing">
          <div class="box-wrapper">
            <Checkbox><input on:click={() => toggleFriendSelection(friend)} type="checkbox" /></Checkbox>
          </div>
        </svelte:fragment>
      </ListItemLabel>
    {/each}
  </ul>
</div>

</div>

<style>
  .spacer {
    height: 1rem;
  }
  .container 
  {
    display:flex;
    justify-content: center;

  }
  .friend-img {
    height:2rem;
  }
  .create-group-container {
    display: flex;
    flex-direction: column;
    max-width: 40rem;
  }
  .button-group {
    display: flex;
    gap: 1rem;
  }
  .create-group-create-group-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  .friends-menu {
    margin-left: 2rem;
  }
  .friends-menu ul {
    list-style-type: none;
    padding: 0;
  }
  .friends-menu li {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 10px;
  }
  .friends-menu button {
    background-color: #007bff;
    color: white;
    border: none;
    padding: 5px 10px;
    font-size: 14px;
    cursor: pointer;
    border-radius: 5px;
  }
  .friends-menu button:hover {
    background-color: #0056b3;
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
  .avatar {
    display: flex;
    height: 200px;
    width: 200px;
  }
</style>