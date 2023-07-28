<script lang="ts">
    import Dragable from "./Dragable.svelte";
</script>

<div>
    
    <Dragable>
        <h2>Berechnung</h2>
        <label for="value1">Wert1</label>
        <input type="text" name="value1"/>        
        

        <label for="operator">Rechen Operation</label>
        <input type="text" name="operator"/>        

        <label for="value2">Wert2</label>
        <input type="text" name="value2"/>        
        <br>
        <input type="radio"/>
        <input type="radio"/>
    </Dragable>

</div>
