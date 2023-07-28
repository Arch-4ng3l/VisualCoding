<script>
    import Dragable from "./Dragable.svelte";
</script>

<style>
</style>

<Dragable>
    <div>
        <h2>If</h2> 
        <label for="condition1">Wert 1</label>

        <input type="text" name="condition1"/>
        <label for="operator">Operator</label>
        <input type="text" name="operator"/>
        <br>
        <label for="condition2">Wert 2</label>
        <input type="text" name="condition2"/>
        <br>
        <br>
        <input type="checkbox" name="true"/>
        <input type="checkbox" name="false" value="False"/>
    </div>
</Dragable>
