<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import type { TransitionConfig } from "svelte/transition";
  import { easeEmphasizedAccel, easeEmphasizedDecel } from "$lib/misc/easing";
  import { outroClass } from "$lib/misc/animation";
</script>

<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
<div class="m3-container">
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <slot />
</div>

<style>
  :root {
    --m3-bottom-sheet-shape: var(--m3-util-rounding-extra-large);
  }

  .m3-container {
    top: 1.5rem;
    position: sticky;
    margin-bottom: 0;
    width: 30%;
    height: calc(100vh - 3rem);
    max-width: 30rem;
    overflow: hidden;
    touch-action: none;

    background-color: rgb(var(--m3-scheme-surface-container-low));
    color: rgb(var(--m3-scheme-on-surface));
    border-radius: var(--m3-bottom-sheet-shape);
    border: none;
    padding: 0;
  }
  .container {
  }
  @keyframes backdrop {
    0% {
      background-color: transparent;
    }
    100% {
      background-color: rgb(var(--m3-scheme-scrim) / 0.5);
    }
  }
  @keyframes backdropReverse {
    0% {
      background-color: rgb(var(--m3-scheme-scrim) / 0.5);
    }
    100% {
      background-color: transparent;
    }
  }
  @media (forced-colors: active) {
    .handle {
      background-color: canvastext;
    }
  }
</style>
