<script>
    import { Button } from '$lib';

export let data;
</script>


<h1 class="m3-font-display-large">Your groups</h1>
{#each data.groups as g}
       <a href="g/{g.id}"><Button type="filled">{g.name}</Button></a>
{/each}
