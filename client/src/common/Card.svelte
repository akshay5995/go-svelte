<script>
  export let heading = "place holder";
</script>

<style>

</style>

<div class="text-center p-2 shadow-lg bg-white my-6 mx-6 rounded-lg">
  <h2 class="font-bold border-teal-400 border-solid border-2 rounded-lg p-3">
    {heading}
  </h2>
  <div class="px-5">
    <slot />
  </div>
</div>
