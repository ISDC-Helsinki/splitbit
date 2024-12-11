<script lang="ts">
    import Button from "$lib/buttons/Button.svelte";
    import Icon from "$lib/misc/_icon.svelte";
    import FamilyRestroom from "@ktibow/iconset-material-symbols/family-restroom";
    import Unarchive from "@ktibow/iconset-material-symbols/unarchive"
    import SyncAlt from "@ktibow/iconset-material-symbols/sync-alt"
    import Group from "@ktibow/iconset-material-symbols/group"
    export let name: String;
    export let debtAmount = 0; 
    export let id;
    export let date;
    
    function handleUnarchive() {
        
        const confirmation = window.confirm(`Do you really want to unarchive ${name}?`);
        if (confirmation) {
            console.log(`Group {id} unarchived.`);
        } else {
            console.log("Unarchiving canceled.");
        }

    }
</script>

<div class="card">
    <div class="left">
        <div>
            <div class="row">
                <div>
                    <Icon icon={FamilyRestroom} class="icon-group-card"/>
                </div>

                <div>
                    <div>
                        Archived on {date}
                    </div>

                    <div class="name">
                        {name}
                    </div>
                </div>
            </div>
            
        </div>

        <div class="row">
            <div>
                <Icon icon={SyncAlt}/>
            </div>

            <div class="debt">
                You owe: ${debtAmount}
            </div>
        </div>

        <div class="row">
            <div>
                <Icon icon={Group}/>
            </div>

            <div class="debt">
                3 members: Member1, Member2, Member3
            </div>
        </div>
    </div>

    

    <div class="right">
        <div>
            <Button type="outlined" on:click={handleUnarchive}>
                <div class="row">
                    <div>
                        <Icon icon = {Unarchive}/>
                    </div> 

                    <div>
                        Unarchive
                    </div>
                </div>
            </Button>
        </div>
        
        <div>
            <a href="g/{id}"><Button type="filled">View</Button></a>
        </div>
    </div>
</div>

<style>
    .card {
        width: 35vw;
        height: 30%; 
        border-radius: 8px;
        overflow: hidden;
        display: flex;
        flex-direction: row;
        background: linear-gradient(to left, rgba(0, 0, 0, 0) -10%, rgba(10, 10, 10, 10) 100%), url("/src/poolcover.jpg");
        background-size: cover;
        background-position: center;
        justify-content: space-between;
        padding: 7px;
        color: whitesmoke;
    }
    
    .left {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
    }

:global(.icon-group-card) {
    display: inline-flex;
    justify-content: center;
    align-items: center;
    border-radius: 50%;
    padding: 8px;
    font-size: 24px;
    color: rgb(41, 51, 51);
    }

    .name {
        font-size: 1.3rem; 
        font-weight: bold; 
        margin: 0;
    }

    .debt {
        padding: 0px;
    }

    .right {
        display: flex;
        flex-direction: column;
        align-items: flex-end;
        gap: 7px;
    }

    .row {
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 10px;
    }

</style>
