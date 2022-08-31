export const fetchCountsByNoun = async (query: string) =>
  await useAsyncData(
    "get",
    () =>
      $fetch("/api/countByNoun", {
        params: {
          query,
        },
      }),
    { initialCache: false }
  );
