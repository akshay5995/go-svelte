<script>
  import Icon from "fa-svelte";

  let github = "https://github.com"
  let siteTitle= client ? client.title : "Generic Title";
  let blogLink = client ? client.blog : "#";
  let githubUser = client ? client.user : '';
  let githubLink = `${github}/${githubUser}`;
</script>

<style>

</style>

<nav class="flex items-center justify-between flex-wrap bg-teal-500 p-6">
  <div class="flex items-center flex-shrink-0 text-white mr-6">
    <a href="/" class="font-semibold text-xl tracking-tight mx-2">
      {siteTitle}
    </a>
  </div>
  <div
    class="w-full block flex-grow lg:flex lg:items-center lg:w-auto text-right">
    <div class="text-sm lg mx-1">
      <a
        href={blogLink}
        class="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white">
        Blog
      </a>
    </div>
    <div class="text-sm lg mx-1">
      <a
        href={githubLink}
        class="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white">
        Github
      </a>
    </div>
  </div>
</nav>
