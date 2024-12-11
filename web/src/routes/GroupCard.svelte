<script lang="ts">
    import Button from "$lib/buttons/Button.svelte";
    import Icon from "$lib/misc/_icon.svelte";
    import FamilyRestroom from "@ktibow/iconset-material-symbols/family-restroom";
    import Archive from "@ktibow/iconset-material-symbols/archive"
    import SyncAlt from "@ktibow/iconset-material-symbols/sync-alt"
    import Group from "@ktibow/iconset-material-symbols/group"
    export let name: String;
    export let debtAmount = 0; 
    export let id;
    
    function handleArchive() {
        
        const confirmation = window.confirm(`Do you really want to archive ${name}?`);
        if (confirmation) {
            console.log(`Group {id} archived.`);
        } else {
            console.log("Archiving canceled.");
        }

    }
</script>

<div class="card">
    <div class="top">
        <div class="content">
            <div class="row">
                <div style="align-items: center;">
                    <Icon icon={FamilyRestroom} class="icon-group-card"/>
                </div>

                <div class="name">
                    {name}
                </div> 
            </div>

            <div class="row">
                <div>
                    <Icon icon={SyncAlt}/>
                </div>

                <div>
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
    </div>



    <div class="bottom-buttons">
        <div>
            <a href="g/{id}"><Button type="filled">View</Button></a>
        </div>

        <div>
            <div>
                <Button type="outlined" on:click={handleArchive}>
                    <div class="row">
                        <div>
                            <Icon icon={Archive}/>
                        </div>
                    
                        <div>
                            Archive
                        </div>
                    </div>
                </Button>
            </div>
        </div>
    </div>
</div>

<style>
    .card {
        width: 30%;
        max-width: 380px;
        height: 50vw; 
        max-height: 400px;
        border-radius: 8px;
        overflow: hidden;
        display: flex;
        flex-direction: column;
        background: linear-gradient(to bottom, rgba(0, 0, 0, 0) -10%, rgba(10, 10, 10, 10) 100%), url("/src/poolcover.jpg");
        background-size: cover;
        background-position: center;
        color: white; 
        justify-content: flex-end;
    }
    
    .top {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        padding: 10px;
    }

:global(.icon-group-card) {
    display: inline-flex;
    justify-content: flex-end ;
    background-color: whitesmoke; 
    border-radius: 50%; 
    padding: 8px;
    }

    .content {
        margin: 0;
        display: flex;
        flex-direction: column;
        gap: 15px;
    }

    .name {
        font-size: 1.5rem; 
        font-weight: bold; 
        align-items: center;
    }

    .bottom-buttons {
        display: flex;
        flex-direction: row-reverse;
        gap: 10px;
        padding: 10px;
    }

    .row {
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 10px;
    }
</style>
