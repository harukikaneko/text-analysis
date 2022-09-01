import { CountsByNoun } from "~~/types/noun";

const ctx = useRuntimeConfig();

export default defineEventHandler(async (event) => {
  const params = useQuery(event);
  const result: CountsByNoun[] = await $fetch("/nouns", {
    baseURL: ctx.baseURL,
    params,
  });

  return result;
});
