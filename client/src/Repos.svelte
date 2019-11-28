<script>
  import { onMount } from "svelte";
  import { faStar } from "@fortawesome/free-solid-svg-icons/faStar";
  import { faCog } from "@fortawesome/free-solid-svg-icons/faCog";
  import { faCodeBranch } from "@fortawesome/free-solid-svg-icons/faCodeBranch";
  import { faCheckCircle } from "@fortawesome/free-solid-svg-icons/faCheckCircle";

  import Card from "./Card.svelte";

  import Icon from "fa-svelte";

  let repoList = [];

  let count = 5;

  let loading = false;

  let isLoadedAll = false;

  async function getRepos() {
    const res = await fetch(`http://localhost:4001/repos?count=${count}`);
    const { repos, loadedAll } = await res.json();
    repoList = repos;
    isLoadedAll = loadedAll;
    loading = false;
  }

  function handleClick() {
    loading = true;
    count += 5;
    getRepos();
  }

  onMount(getRepos);
</script>

<style>

</style>

<Card title="Top repos">
  <div class="px-5">
    {#if repoList.length}
      <div>
        {#each repoList as { full_name, html_url, stargazers_count, forks_count }}
          <div class="flex flex-wrap m-2">
            <a class="font-semibold w-1/2 flex-grow text-left" href={html_url}>
              {full_name}
            </a>
            <p class="mx-2 flex leading-none p-1">
              {#if stargazers_count > 0}{stargazers_count}{:else}-{/if}
              <Icon
                class="h-4 mx-1 fill-current text-yellow-500"
                icon={faStar} />
              {#if forks_count > 0}{forks_count}{:else}-{/if}
              <Icon class="h-4 mx-1 fill-current" icon={faCodeBranch} />
            </p>
          </div>
        {/each}
        {#if !isLoadedAll}
          <button
            class="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold
            hover:text-white p-1 px-2 border border-blue-500
            hover:border-transparent rounded h-10 my-2"
            on:click={handleClick}>
            {#if loading}
              <Icon class="h-4 mx-1 fill-current" icon={faCog} />
            {:else}Load more{/if}
          </button>
        {:else}
          <p class="w-41 flex leading-none p-4 font-bold">
            <Icon
              class="h-4 mx-1 fill-current text-green-500"
              icon={faCheckCircle} />
            Loaded all public repos
          </p>
        {/if}
      </div>
    {:else}
      <h2 class="font-semibold text-grey-500 cursor-wait">
        <div class="sp sp-bars" />
      </h2>
    {/if}
  </div>
</Card>
