<script>
  import Card from "./common/Card.svelte";
  import { onMount } from "svelte";

  let name, email, bio, location, image;

  let loading = true;

  async function getUser() {
    const resp = await fetch("/user");
    const data = await resp.json();

    name = data.name;
    email = data.email;
    bio = data.bio;
    location = data.location;
    image = data.avatar_url;

    loading = false;
  }

  onMount(getUser);
</script>

<style>

</style>

<Card heading="About">
  {#if loading}
    <h2 class="font-semibold text-grey-500 cursor-wait">
      <div class="sp sp-bars" />
    </h2>
  {:else}
    <div class="bg-white rounded-lg p-6">
      <img
        class="h-16 w-16 rounded-full mx-auto"
        alt="profile picture"
        src={image} />
      <div>
        <h2 class="text-lg font-semibold">{name}</h2>
        <div class="text-purple">{bio}</div>
        <a href={`mailto:${email}`}>{email}</a>
        <div>{location}</div>
      </div>
    </div>
  {/if}
</Card>
