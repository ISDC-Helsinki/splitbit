<script lang="ts">
    import Button from "$lib/buttons/Button.svelte";
    import Icon from "$lib/misc/_icon.svelte";
    import FamilyRestroom from "@ktibow/iconset-material-symbols/family-restroom";
    import LocationOn from "@ktibow/iconset-material-symbols/location-on";
    import Description from "@ktibow/iconset-material-symbols/description";
    import Image from "@ktibow/iconset-material-symbols/image";
    import ShoppingCart from "@ktibow/iconset-material-symbols/shopping-cart";
    export let name: String;
    export let debtAmount = 0; 
    export let expense = "itemname"
</script>

<div class="card-holder">
    <div class="left">
        <div class="icons">
            <Icon icon={FamilyRestroom}/>
        </div>
            
        <div class="column-in-group">
            <div>
                In group
            </div>
                
            <div>
                <b>{name}</b>
            </div>
        </div>

        <div class="line"></div>

        <div class="centered">
            Date
        </div>

        <div class="icons">
            <Icon icon={ShoppingCart}/>
        </div>

        <div class="column">
            <div>
                {expense}
            </div>

            <div>
                Mariusz paid {debtAmount}
            </div>
        </div>

    </div>

    <div class="right">
        <div class="icons">
            <Icon icon={LocationOn}/>
        </div>

        <div class="icons">
            <Icon icon={Description}/>
        </div>

        <div class="icons">
            <Icon icon={Image}/>
        </div>

        <div class="column">
            <div class="red-font">
                You lent
            </div>

            <div class="red-font">
                {debtAmount}$
            </div>
        </div>

        <div>
            <Button type="tonal">View in group</Button>
        </div>

    </div>
</div>

<style>
    .card-holder {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
    }

    .left {
        display: flex;
        flex-direction: row;
        gap: 20px;
    }

    .right {
        display: flex;
        flex-direction: row;
        justify-items: flex-end;
        align-items: center;
        gap: 20px;
    }

    .column {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .column-in-group {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 100px;
    }

    .icons {
        display: flex;
        align-items: center;
        font-size: 30px;
    }

    .centered {
        display: flex;
        align-items: center;
    }

    .line {
        border-left: 2px solid darkslategrey;
        height: 100%;
        border-radius: 1%;
    }

    .red-font {
        color: rgb(250, 18, 18);
    }
</style>