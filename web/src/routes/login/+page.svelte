<script>
    import { goto } from '$app/navigation';
    import { Button, TextFieldOutlined, Card } from '$lib';
    import CalendarPicker from '$lib/forms/_picker/CalendarPicker.svelte';
  
    let instanceUrl = '';
    let username = '';
    let password = '';
  
    async function handleLogin() {
      try {
        const response = await fetch(`${instanceUrl}/login`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ username, password })
        });
  
        if (response.ok) {
          const data = await response.json();
          // Save the token or user data as needed
          console.log('Login successful:', data);
          // Redirect to the main page or dashboard
          goto('/');
        } else {
          console.error('Login failed:', response.statusText);
        }
      } catch (error) {
        console.error('Error during login:', error);
      }
    }
  </script>
  
  <form on:submit|preventDefault={handleLogin}>
    <center>
    <Card type="elevated" display="square">
    <img src="slogo.svg" alt="Logo" width="90px" height="90px" >
    <h1>Welcome to SplitBit Instance</h1>
    <div>
      <p class="m3-font-body-medium"> Hosted at splitbit.isdc.fi </p>
    </div>
    <div>
      <TextFieldOutlined name = "username" bind:value={username} required />
    </div>
    <div>
        <TextFieldOutlined name = "password" bind:value={password} required />
      </div>
    <Button type="filled">Log in</Button>
    </Card>
  </center>
  </form>