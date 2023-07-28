<script>
    let count = 0;
    function handleClick() {
        count += 1;
    }
</script>

<style>

</style>

<main>
    <div>
        <button on:click={handleClick}>
            clicked {count}
    </button>
    </div>
</main>
