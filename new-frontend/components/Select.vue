<script setup lang="ts">
  import { useJobtitlesStore } from "../store/jobtitles";
  import { useCompaniesStore } from "../store/companies";
  import { useCitiesStore } from "../store/cities";
  import { useSenioritiesStore } from "../store/senorities";
  import { useRatingsStore } from "../store/ratings";

  const jobtitleStore = useJobtitlesStore();
  const companyStore = useCompaniesStore();
  const cityStore = useCitiesStore();
  const seniorityStore = useSenioritiesStore();
  const ratingStore = useRatingsStore();

  const myfilterjob: Ref<string> = ref('');
  const myfiltercompany: Ref<string> = ref('');
  const myfilterseniority: Ref<string> = ref('');
  const myfiltercity: Ref<string> = ref('');

  const companies = computed(() => companyStore.items);
  const cities = computed(() => cityStore.items);
  const jobtitles = computed(() => jobtitleStore.items);
  const seniorities = computed(() => seniorityStore.items);
  const filterjob = computed(() => jobtitleStore.filter);
  const filtercompany = computed(() => companyStore.filtercompany);
  const filterseniority = computed(() => seniorityStore.filter);
  const filtercity = computed(() => cityStore.filter);
  const page = computed(() => ratingStore.page);
  const limit = computed(() => ratingStore.limit);
  
  onMounted(async () => {
    await fetchCompanies();
    await fetchCities();
    await fetchJobtitles();
    await fetchSeniorities();
  });

  const onChangeJobTitle = async () => {
    jobtitleStore.filter = myfilterjob.value;
    await reloadData();
  }
  
  const onChangeSeniority = async () => {
    seniorityStore.filter = myfilterseniority.value;
    await reloadData();
  }

  const onChangeCompany = async () => {
    companyStore.filtercompany = myfiltercompany.value;
    await reloadData();
  }
  
  const onChangeCity = async () => {
    cityStore.filter = myfiltercity.value;
    await reloadData();
  }

  const reloadData = async () => {
    await ratingStore.getItems({
      company: filtercompany.value,
      jobtitle: filterjob.value,
      seniority: filterseniority.value,
      city: filtercity.value,
    });
    await ratingStore.getAverage({
      company: filtercompany.value,
      jobtitle: filterjob.value,
      seniority: filterseniority.value,
      city: filtercity.value,
    });
    ratingStore.page = 1;
  }

  const fetchCompanies = async () => {
    await companyStore.getItems();
  }

  const fetchCities = async () => {
    await cityStore.getItems();
  }

  const fetchJobtitles = async () => {
    await jobtitleStore.getItems();
  }

  const fetchSeniorities = async () => {
    await seniorityStore.getItems();
  }
</script>

<template>
  <div class="flex justify-center items-center flex-col sm:flex-row">
    <div class="mb-3 w-full sm:w-1/3 lg:w-40 ml-0 lg:ml-6">
      <h4
        style="color: #b1b1b1"
        class="pb-1 font-semibold text-xs md:text-base"
      >
        Seniority
      </h4>
      <select
        class="form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
        aria-label="Default select example"
        v-model="myfilterseniority"
        @change="onChangeSeniority"
      >
        <option
          class="text-xs md:text-sm"
          value=""
          disabled
          style="font-family: 'Inter', sans-serif"
        >
          All Seniorities
        </option>
        <option
          class="text-xs md:text-sm"
          value=""
          style="font-family: 'Inter', sans-serif"
        >
          any
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          v-for="(seniority, index) in seniorities"
          :key="index"
          :value="seniority"
          v-show="seniority !== ''"
          class="text-xs md:text-sm"
        >
          {{ seniority }}
        </option>
      </select>
    </div>
    <div class="mb-3 w-full sm:w-1/3 lg:w-40 ml-0 lg:ml-6">
      <h4
        style="color: #b1b1b1"
        class="pb-1 font-semibold text-xs md:text-base"
      >
        Job Title
      </h4>
      <select
        class="form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
        aria-label="Default select example"
        v-model="myfilterjob"
        @change="onChangeJobTitle"
      >
        <option
          class="text-xs md:text-sm"
          value=""
          disabled
          style="font-family: 'Inter', sans-serif"
        >
          All Titles
        </option>
        <option
          class="text-xs md:text-sm"
          value=""
          style="font-family: 'Inter', sans-serif"
        >
          any
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          v-for="(link, index) in jobtitles"
          :key="index"
          :value="link"
          v-show="link !== ''"
          class="text-xs md:text-sm"
        >
          {{ link }}
        </option>
      </select>
    </div>
    <div class="mb-3 w-full sm:w-1/3 lg:w-40 ml-0 sm:ml-6">
      <h4
        style="color: #b1b1b1"
        class="pb-1 font-semibold text-xs md:text-base"
      >
        Companies
      </h4>
      <select
        class="form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
        aria-label="Default select example"
        v-model="myfiltercompany"
        @change="onChangeCompany"
      >
        <option
          class="text-xs md:text-sm"
          value=""
          disabled
          style="font-family: 'Inter', sans-serif"
        >
          All Compagnies
        </option>
        <option
          class="text-xs md:text-sm"
          value=""
          style="font-family: 'Inter', sans-serif"
        >
          any
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          v-for="link in companies"
          :key="link.id"
          :value="link.name"
          v-show="link.name !== ''"
          class="text-xs md:text-sm"
        >
          {{ link.name }}
        </option>
      </select>
    </div>
    <div class="mb-3 w-full sm:w-1/3 lg:w-40 ml-0 sm:ml-6">
      <h4
        style="color: #b1b1b1"
        class="pb-1 font-semibold text-xs md:text-base"
      >
        Cities
      </h4>
      <select
        class="form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
        aria-label="Default select example"
        v-model="myfiltercity"
        @change="onChangeCity"
      >
        <option
          class="text-xs md:text-sm"
          value=""
          disabled
          style="font-family: 'Inter', sans-serif"
        >
          All Cities
        </option>
        <option
          class="text-xs md:text-sm"
          value=""
          style="font-family: 'Inter', sans-serif"
        >
          any
        </option>
        <option
          style="font-family: 'Inter', sans-serif"
          v-for="(city, index) in cities"
          :key="index"
          :value="city"
          v-show="city !== ''"
          class="text-xs md:text-sm"
        >
          {{ city }}
        </option>
      </select>
    </div>
  </div>
</template>