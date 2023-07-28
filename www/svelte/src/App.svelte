<script>
    
    import {Router, Link, Route} from 'svelte-routing';
    import Test from './Test.svelte';
    import Dashboard from './Dashboard.svelte';

</script>


<Router>

    <nav>
    <Link to="/test">Test</Link>
    <Link to="/dashboard">Test</Link>
    </nav>

    <Route path="/dashboard" component={Dashboard}/>
    <Route path="/test" component={Test}/>
</Router>

