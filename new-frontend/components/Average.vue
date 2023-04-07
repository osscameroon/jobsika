<script setup lang="ts">
  import { useRatingsStore } from "../store/ratings";
  import { useJobtitlesStore } from "../store/jobtitles";
  import { useCompaniesStore } from "../store/companies";

  const { title } = defineProps({
		title: {type: String, required: true}
	});

  const ratingStore = useRatingsStore();
  const jobtitleStore = useJobtitlesStore();
  const companyStore = useCompaniesStore();
  const filterjob = computed(() => jobtitleStore.filter);
  const filtercompany = computed(() => companyStore.filtercompany);
  const average = computed(() => ratingStore.average);
  const stars = computed(() => ratingStore.averageStars);

  const starsPicture: Ref<string> = ref('/star.png');
  const salaryRating: Ref<string> = ref('');
  const errored: Ref<boolean> = ref(false);

  onMounted(async () => {
    await fetchAverage();
  })

  const fetchAverage = async () => {
    await ratingStore.getAverage({
      company: filtercompany.value,
      jobtitle: filterjob.value,
    })
  }
</script>

<template>
  <div>
    <div v-if="errored" class="text-center py-4 md:py-10 font-bold text-sm" />
    <div
      v-else
      style="background: #235365; font-family: 'Inter', sans-serif"
      class="site__average py-2 md:py-4 opacity-60 rounded-md my-3"
    >
      <div class="flex flex-row justify-between items-center">
        <div class="site__average__title w-full md:w-1/6 ml-10">
          <p
            class="text-white text-xs lg:text-sm font-medium"
            style="font-family: 'Inter', sans-serif"
          >
            {{ title }}
          </p>
        </div>
        <div
          class="
            site__average__content
            w-full
            flex flex-col
            md:flex-row
            items-center
            md:w-1/2
            ml-10
            mr-2
            md:ml-0 md:mr-0
          "
        >
          <p
            v-if="average !== 0"
            class="
              text-white text-xs
              lg:text-sm
              font-medium
              md:ml-20
              xl:pl-6
              text-center
            "
            style="font-family: 'Inter', sans-serif"
          >
            {{ average }}
          </p>
          <p v-else class="text-white text-xs lg:text-sm text-center">loading..</p>
          <div class="flex pt-1 md:pt-0 ml-0 md:ml-8 xl:ml-20 md:mr-10">
            <div v-for="(item, index) in stars" :key="index" class="flex">
              <img class="w-4 h-4 mr-1" :src="starsPicture" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
