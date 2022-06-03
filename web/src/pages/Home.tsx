import purify from "dompurify";
import { marked } from "marked";
import { nanoid } from "nanoid";
import { onMount } from "solid-js";

import { BasePage } from "../components/BasePage";

const contentId = `content-home-${nanoid()}`;

export const Home = () => {
  onMount(async () => {
    const content = await (await fetch("/content/home.md")).text();
    const md = purify.sanitize(marked.parse(content));
    const inner = document.getElementById(contentId);

    inner.innerHTML = md;

    inner
      .querySelectorAll("a")
      .forEach((item) => item.classList.add("text-blue-700"));

    inner.querySelectorAll("h1").forEach((item) => {
      item.classList.add("py-2");
      item.classList.add("text-4xl");
      item.classList.add("font-400");
    });

    inner.querySelectorAll("h2").forEach((item) => {
      item.classList.add("py-1");
      item.classList.add("text-2xl");
      item.classList.add("font-400");
    });
  });

  return (
    <BasePage>
      <div class="flex flex-col gap-y-2" id={contentId} />
    </BasePage>
  );
};

export default Home;
