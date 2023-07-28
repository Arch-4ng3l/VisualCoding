<script>
    
    import If from './If.svelte';
    import Calc from './Calc.svelte';
    import { writable } from 'svelte/store';
    let spawnedObjects = writable([]);
    let i = 0;

    function handleSpawn(type) {
        spawnedObjects.update(arr=> [...arr, {type: type}]);
    }

</script>

<main>
    <h1>Hello</h1>
    <button on:click={() => handleSpawn("if")}>If</button>
    <button on:click={() => handleSpawn("calc")}>Berechnung</button>

    {#each $spawnedObjects as {type}}
        {#if  type == "if"}
        <If id={i}/>
        {:else if type == "calc"}
        <Calc id={i}/>
        {/if}
    {/each}
</main>


<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
