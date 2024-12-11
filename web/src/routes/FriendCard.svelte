<script lang="ts">
    export let name: String;
    export let debtAmount = 0; 
</script>

<div class="card">
    <div class="left">
        <div class="profile-picture">
            <img src="/src/poolcover.jpg" alt="profilepicture">
        </div>
    </div>

    <div class="right">
        <div>
            <b>{name}</b>
        </div>

        <div>
            Part of 2 groups, owes you {debtAmount}
        </div>
    </div>
</div>

<style>
    .card {
        width: 60vw;
        height: 30%; 
        overflow: hidden;
        display: flex;
        flex-direction: row;
        gap: 15px;
    }
    
    .left {
        display: flex;
        flex-direction: column;
    }

    .profile-picture img {
    width: 60px;
    height: 60px;
    border-radius: 50%; 
    object-fit: cover;
    }

    

    .right {
        display: flex;
        flex-direction: column;
        justify-content: center;
    }
</style>
