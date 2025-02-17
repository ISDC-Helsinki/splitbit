<script>
  import { goto } from "$app/navigation";
  import { Button, TextFieldOutlined, Card } from "$lib";
  import api from "$lib/api/api";
  import CalendarPicker from "$lib/forms/_picker/CalendarPicker.svelte";

  /**
   * @param {string} str
   */
  function base64UrlDecode(str) {
    // Replace URL-safe characters
    str = str.replace(/-/g, "+").replace(/_/g, "/");
    // Pad with '=' to make it a valid Base64 string
    while (str.length % 4) {
      str += "=";
    }
    return JSON.parse(atob(str));
  }

  let username = "";
  let password = "";

  async function handleLogin() {
    try {
      const {
        data, // only present if 2XX response
        error, // only present if 4XX or 5XX response
        response,
      } = await api.POST("/login", { body: { username, password } });
      if (response.ok && data != undefined) {
        // Save the token or user data as needed
        console.log("Login successful:", data.token);
        // Split the JWT into its components
        const parts = data.token.split(".");
        if (parts.length !== 3) {
          throw new Error("Invalid JWT token");
        }

        // Decode the payload
        const payload = base64UrlDecode(parts[1]);

        // Extract username and userID
        const username = payload.username;
        const userID = payload.id;

        // Store in localStorage
        localStorage.setItem("username", username);
        localStorage.setItem("userID", userID);

        // Optional: Log the values to the console
        console.log("Username:", username);
        console.log("User ID:", userID);
        // Redirect to the main page or dashboard
        goto("/");
      } else {
        console.error("Login failed:", response.statusText);
      }
    } catch (error) {
      console.error("Error during login:", error);
    }
  }
</script>

<main>
  <form class="elevation-3" on:submit|preventDefault={handleLogin}>
    <img src="slogo.svg" alt="Logo" width="90px" height="90px" />
    <h1>Welcome to SplitBit Instance</h1>
    <div>
      <p class="m3-font-body-medium">Hosted at splitbit.isdc.fi</p>
    </div>
    <div>
      <TextFieldOutlined name="username" bind:value={username} required />
    </div>
    <div>
      <TextFieldOutlined
        extraOptions={{ type: "password" }}
        name="password"
        bind:value={password}
        required
      />
    </div>
    <Button type="filled">Log in</Button>
  </form>
</main>

<style lang="scss">
  form {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: fit-content;
    margin: auto;
    padding: 3rem 4rem;
    height: fit-content;
    justify-content: center;
    border-radius: 0.8rem;
    background-color: rgb(var(--m3-scheme-surface-container-lowest));
    div {
      margin-bottom: 0.5rem;
    }
  }
  h1 {
    text-align: center;
  }
  main {
    height: 100%;
    display: flex;
  }
  @media (prefers-color-scheme: light) {
    img {
      filter: invert(1);
    }
  }
</style>
