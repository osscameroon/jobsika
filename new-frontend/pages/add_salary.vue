<script setup lang="ts">
	import Button from "../components/Button.vue";
	import Notification from "../components/Notification.vue";
	import { ICompany } from "../types/company";
	import { IRating } from "../types/ranting";
	import { ITooltip } from "~/types/common";
	import { useRatingsStore } from "../store/ratings";
  import { useJobtitlesStore } from "../store/jobtitles";
  import { useCompaniesStore } from "../store/companies";
	import { useCitiesStore } from "../store/cities";
	import { useSenioritiesStore } from "../store/senorities";

	const DEFAULT_RATING: IRating = {
		company_name: '', salary: 0, city: '', seniority: '', rating: 0,
		comment: '', job_title: '', createdat: '', salary_id: 0,
		company_id: 0, company_rating_id: 0, country: ''
	}

	const router = useRouter();

	const ratingStore = useRatingsStore();
	const jobtitleStore = useJobtitlesStore();
	const companyStore = useCompaniesStore();
	const cityStore = useCitiesStore();
	const seniorityStore = useSenioritiesStore();

	const isOpen: Ref<boolean> = ref(false);
	const errorSalary: Ref<string> = ref('');
	const errorJobtitle: Ref<string> = ref('');
	const errorCity: Ref<string> = ref('');
	const errorAlert: Ref<string> = ref('');
	const starsPicture: Ref<string> = ref('/star.png');
	const warning: Ref<string> = ref('/warning.png');
	const opened: Ref<any[]> = ref([]);
	const newRating: Ref<IRating> = ref(DEFAULT_RATING);
	const isCompanyNameFocus: Ref<boolean> = ref(false);
	const isJobTitleFocus: Ref<boolean> = ref(false);
	const isCityFocus: Ref<boolean> = ref(false);
	const tooltips: Ref<ITooltip[]> = ref([
		{
			id: 1,
			name: 'company',
		},
		{
			id: 2,
			name: 'monthly',
		},
		{
			id: 3,
			name: 'comment',
		},
		{
			id: 4,
			name: 'jobtitle',
		},
		{
			id: 5,
			name: 'city',
		},
		{
			id: 6,
			name: 'seniority',
		},
		{
			id: 7,
			name: 'rate',
		}
	]);

	const companies = computed(() => companyStore.items);
	const jobtitles = computed(() => jobtitleStore.items);
	const cities = computed(() => cityStore.items);
	const seniorities = computed(() => seniorityStore.items);
	const companyNames = computed(() => companies.value.map((elem: ICompany) => elem.name));
	const filteredCompanyNames = computed(() => {
		if (newRating.value.company_name === '') {
			return []
		} else {
			return companyNames.value.filter((elem: string) => {
				return !elem
				.toLowerCase()
				.indexOf(newRating.value.company_name.toLowerCase())
			})
		}
	});
	const filteredJobTitles = computed(() => {
		if (newRating.value.job_title === '') {
			return []
		} else {
			return jobtitles.value.filter((elem: string) => {
				return !elem
					.toLowerCase()
					.indexOf(newRating.value.job_title.toLowerCase())
			})
		}
	});
	const filteredCities = computed(() => {
		if (newRating.value.city === '') {
			return []
		} else {
			return cities.value.filter((elem) => {
				return !elem.toLowerCase().indexOf(newRating.value.city.toLowerCase())
			})
		}
	});
	const companyComplation = computed(() => {
		return filteredCompanyNames.value.length > 0 && isCompanyNameFocus.value
	});
	const jobTitleComplation = computed(() => {
		return filteredJobTitles.value.length > 0 && isJobTitleFocus.value
	});
	const cityComplation = computed(() => {
		return filteredCities.value.length > 0 && isCityFocus.value
	});

	const fSalaryValue = computed({
		get: () => {
			if (newRating.value.salary) {
				return formatUSD(newRating.value.salary);
			}
		},
		set: (newValue: string | any) => {
			newRating.value.salary = parseUSD(newValue);
		}
	});

	onMounted(async () => {
		await fetchCompanies();
		await fetchJobtitles();
		await fetchCities();
		await fetchSeniorities();
	});

	// METHODS
	const goToList = () => {
		router.push('/')
	}

	const setCompanyName = (name: string) => {
		newRating.value.company_name = name;
		isCompanyNameFocus.value = false;
	}

	const setJobTitle = (jobtitle: string) => {
		newRating.value.job_title = jobtitle;
		isJobTitleFocus.value = false;
	}

	const setCity = (city: string) => {
		newRating.value.city = city;
		isCityFocus.value = false;
	}

	const fetchCompanies = async () => {
		await companyStore.getItems();
	}
	const fetchJobtitles = async () => {
		await jobtitleStore.getItems();
	}

	const fetchCities = async () => {
		await cityStore.getItems();
	}

	const fetchSeniorities = async () => {
		await seniorityStore.getItems();
	}

	const setGrade = (value: number) => {
		newRating.value.rating = value;
	}

	const blurAll = () => {
		isJobTitleFocus.value = false;
		isCompanyNameFocus.value = false;
		isCityFocus.value = false;
	}

	const nameFocus = () => {
		blurAll();
		isCompanyNameFocus.value = true;
	}

	const jobTitleFocus = () => {
		blurAll();
		isJobTitleFocus.value = true;
	}
	const cityFocus = () => {
		blurAll();
		isCityFocus.value = true;
	}

	const addRating = async () => {
		if (String(newRating.value.job_title).length === 0) {
			errorJobtitle.value = 'This field cannot be empty';
			document.getElementById('myAlertJob')?.scrollIntoView({
				behavior: 'smooth',
				block: 'center',
			})
			return
		}
		else if (String(newRating.value.city).length === 0) {
			errorCity.value = 'This field cannot be empty';
			document.getElementById('myAlertCity')?.scrollIntoView({
				behavior: 'smooth',
				block: 'center',
			})
			return
		}
		else if (String(newRating.value.salary).length === 0) {
			errorSalary.value = 'This field cannot be empty';
			errorAlert.value = 'please fill in the compulsory fields';
			document.getElementById('myAlertSalary')?.scrollIntoView({
				behavior: 'smooth',
				block: 'center',
			})
		} else {
			await ratingStore.addItem(newRating.value);
			newRating.value = {...DEFAULT_RATING};
			await router.push('/');
		}
	}

	const onlyNumber = ($event: any) => {
		const keyCode = $event.keyCode ? $event.keyCode : $event.which
		if ((keyCode < 48 || keyCode > 57) && keyCode !== 46) {
			$event.preventDefault()
		}
	}

	const toggle = (id: number) => {
		const index = opened.value.indexOf(id);
		if (index > -1) {
			opened.value.splice(index, 1)
		} else {
			opened.value.push(id)
		}
	}

	const formatUSD = (num: number) => {
		return (
			'' +
			Number(num)
				.toString()
				.replace(/(\d)(?=(\d\d\d)+(?!\d))/g, '$1 ')
		)
	}

	const parseUSD = (text: string) => {
		return Number(text.replace('$', '').replace(/ /g, ''));
	}
</script>

<template>
	<div>
		<main class="site__main pt-6 md:pt-12 pb-10 md:pb-20">
			<div class="container mx-auto w-10/12">
				<Title title="Fill in the form and submit your contribution" font-size="header" only-title="onlyTitle" />
				<p class="text-center text-xs sm:text-base pt-2 md:pt-4 font-normal"
					style="color: #2e2e2e; font-family: 'Inter', sans-serif">
					Your entry is anonymous. Please ensure that it is fair, accurate and
					honest.
				</p>
				<form @submit.prevent="addRating">
					<div class="site__main-row flex flex-col md:w-7/12 mx-auto">
						<div class="site__main-row w-full mt-5 md:mt-16">
							<div class="flex">
								<p class="text-xs font-bold md:text-lg" style="color: #b1b1b1; font-family: 'Inter', sans-serif">
									Company Name
								</p>
								<div class="flex items-center">
									<span
										class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
										:class="{ opened: opened.includes(tooltips[0].id) }" @click="toggle(tooltips[0].id)">
										!
									</span>
								</div>
							</div>
							<div v-if="opened.includes(tooltips[0].id)" class="w-full bg-primary">
								<div class="w-full my-1">
									<p class="font-light text-xs leading-5" style="color: #000000; font-family: 'Inter', sans-serif">
										This text requires you to enter the present or past company
										name you work(ed). Note that if you choose not to fill this
										field, a generic name like "A local company" will be used by
										default. If you choose to enter a company name, it will be
										displayed if it has at least three entries with the same job
										title and town.
									</p>
								</div>
							</div>

							<input :value="newRating.company_name" type="text" style="height: 61px"
								class="site__input-field border border-grayC mt-2 w-full rounded-md mb-4 md:mb-16" @input="
									(event) => (newRating.company_name = (event.target as HTMLInputElement).value)
								" @focus="nameFocus()" />
							<ul v-if="companyComplation" style="background: white"
								class="h-20 md:h-32 overflow-y-scroll rounded-lg -mt-3 md:-mt-14 mb-2 md:mb-4 cursor-pointer">
								<li v-for="(name, index) in filteredCompanyNames" :key="index" class="shadow-sm py-2 px-4"
									@click="setCompanyName(name)">
									{{ name }}
								</li>
							</ul>
							<div class="flex">
								<p class="text-xs md:text-lg font-bold" style="color: #b1b1b1; font-family: 'Inter', sans-serif">
									Job Title
									<span style="color: red">&#42;</span>
								</p>
								<div class="flex items-center">
									<span
										class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
										:class="{ opened: opened.includes(tooltips[0].id) }" @click="toggle(tooltips[3].id)">
										!
									</span>
								</div>
							</div>
							<div v-if="opened.includes(tooltips[3].id)" class="w-full bg-primary">
								<div class="w-full my-1">
									<p class="font-light text-xs leading-5" style="color: #000000; font-family: 'Inter', sans-serif">
										This field requires you to enter the title of the position
										and/or specialization (if applicable). E.g.
										"Accountant(Audit)"
									</p>
								</div>
							</div>
							<input :value="newRating.job_title" type="text" style="height: 61px"
								class="site__input-field border border-grayC mt-2 w-full rounded-md"
								@input="(event) => (newRating.job_title = (event.target as HTMLInputElement).value)" @focus="jobTitleFocus()" />
							<div id="myAlertJob" class="mb-4 md:mb-16 mt-2">
								<notification :message="errorJobtitle" />
							</div>
							<ul v-if="jobTitleComplation" style="background: white"
								class="h-20 md:h-32 overflow-y-scroll rounded-lg -mt-3 md:-mt-14 mb-2 md:mb-4 cursor-pointer">
								<li v-for="(job, index) in filteredJobTitles" :key="index" class="shadow-sm py-2 px-4"
									@click="setJobTitle(job)">
									{{ job }}
								</li>
							</ul>
							<div class="flex">
								<p class="text-xs md:text-lg font-bold" style="color: #b1b1b1; font-family: 'Inter', sans-serif">
									Seniority
								</p>
								<div class="flex items-center">
									<span
										class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
										:class="{ opened: opened.includes(tooltips[0].id) }" @click="toggle(tooltips[5].id)">
										!
									</span>
								</div>
							</div>
							<div v-if="opened.includes(tooltips[5].id)" class="w-full bg-primary">
								<div class="w-full my-1">
									<p class="font-light text-xs leading-5" style="color: #000000; font-family: 'Inter', sans-serif">
										This field requires you to enter the seniority of the job
										title you have/had at the company.
									</p>
								</div>
							</div>
							<select v-model="newRating.seniority"
								class="mt-2 mb-4 md:mb-16 form-select appearance-none block w-full px-3 py-1.5 text-gray-700 bg-white bg-clip-padding bg-no-repeat border border-grayC rounded transition ease-in-out focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none text-xs md:text-sm font-bold"
								style="height: 61px; font-family: 'Inter', sans-serif" aria-label="Default select example"
								@focus="blurAll()">
								<option v-for="seniority in seniorities" :key="seniority" :value="seniority"
									style="font-family: 'Inter', sans-serif" class="text-xs md:text-sm">
									{{ seniority }}
								</option>
							</select>
						</div>
						<div class="flex">
							<p class="text-xs md:text-lg font-bold" style="color: #b1b1b1; font-family: 'Inter', sans-serif">
								City
								<span style="color: red">&#42;</span>
							</p>
							<div class="flex items-center">
								<span
									class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
									:class="{ opened: opened.includes(tooltips[0].id) }" @click="toggle(tooltips[4].id)">
									!
								</span>
							</div>
						</div>
						<div v-if="opened.includes(tooltips[4].id)" class="w-full bg-primary">
							<div class="w-full my-1">
								<p class="font-light text-xs leading-5" style="color: #000000; font-family: 'Inter', sans-serif">
									This field requires you to enter the town you are/were
									employed in. For remote workers, please enter the town you
									are/were based in.
								</p>
							</div>
						</div>

						<input :value="newRating.city" type="text" style="height: 61px"
							class="site__input-field border border-grayC mt-2 w-full rounded-md"
							@input="(event) => (newRating.city = (event.target as HTMLInputElement).value)" @focus="cityFocus()" />
						<div id="myAlertCity" class="mb-4 md:mb-8 mt-2">
							<notification :message="errorCity" />
						</div>
						<ul v-if="cityComplation" style="background: white"
							class="h-20 md:h-32 overflow-y-scroll rounded-lg -mt-3 md:-mt-14 mb-2 md:mb-4 cursor-pointer">
							<li v-for="(city, index) in filteredCities" :key="index" class="shadow-sm py-2 px-4" @click="setCity(city)">
								{{ city }}
							</li>
						</ul>
						<div class="flex">
							<p class="text-xs md:text-lg font-bold" style="color: #b1b1b1; font-family: 'Inter', sans-serif">
								Monthly Salary (FCFA)
								<span style="color: red">&#42;</span>
							</p>
							<div class="flex items-center">
								<span
									class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
									:class="{ opened: opened.includes(tooltips[0].id) }" @click="toggle(tooltips[1].id)">
									!
								</span>
							</div>
						</div>
						<div v-if="opened.includes(tooltips[1].id)" class="w-full bg-primary">
							<div class="w-full my-1">
								<p class="font-light leading-5 text-xs" style="color: #000000; font-family: 'Inter', sans-serif">
									This field requires you to enter the pre-tax/gross salary you
									get, vacation money included. You can find this amount in your
									contract or on your paycheck.
								</p>
							</div>
						</div>
						<input v-model="fSalaryValue" style="height: 61px"
							class="site__input-field border border-grayC mt-2 w-full rounded-md cursor-text pl-2" @keypress="onlyNumber"
							@focus="blurAll()" />
						<div id="myAlertSalary" class="mb-2 md:mb-8 mt-2">
							<notification :message="errorSalary" />
						</div>
						<div class="flex">
							<p class="text-xs md:text-lg font-bold" style="color: #b1b1b1; font-family: 'Inter', sans-serif">
								Rate
							</p>
							<div class="flex items-center">
								<span
									class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
									:class="{ opened: opened.includes(tooltips[0].id) }" @click="toggle(tooltips[6].id)">
									!
								</span>
							</div>
						</div>
						<div v-if="opened.includes(tooltips[6].id)" class="w-full bg-primary">
							<div class="w-full mt-1 mb-2">
								<p class="font-light text-xs leading-5" style="color: #000000; font-family: 'Inter', sans-serif">
									This field requires you to rate your overall experience with
									the company. For example, 1 star could mean you will not
									recommend anybody to work with that company or five stars
									could mean your overall experience with the company was
									excellent and you will gladly recommend anyone to work with
									them.
								</p>
							</div>
						</div>
						<StarRating :grade="newRating.rating" :max-stars="5" :has-counter="true" @changeGrade="setGrade"
							@blurall="blurAll()" />
						<div class="flex">
							<p class="text-xs md:text-lg font-bold" style="color: #b1b1b1; font-family: 'Inter', sans-serif">
								Comments
							</p>
							<div class="flex items-center">
								<span
									class="cursor-pointer h-4 text-center w-4 ml-2 text-grayC rounded-full border border-grayC text-xs flex items-center justify-center"
									:class="{ opened: opened.includes(tooltips[0].id) }" @click="toggle(tooltips[2].id)">
									!
								</span>
							</div>
						</div>
						<div v-if="opened.includes(tooltips[2].id)" class="w-full bg-primary">
							<div class="w-full my-1">
								<p class="font-light text-xs leading-5" style="color: #000000; font-family: 'Inter', sans-serif">
									This field requires you to tell us more about your experience
									in the company. Please share more details about your
									contribution. Don't enter informations that might identify
									you. Any bad or good things that happened that you want to
									share with us?
								</p>
							</div>
						</div>

						<textarea v-model="newRating.comment" type="textarea" rows="5" style="height: 120px"
							class="resize rounded-md w-full h-full mt-2 border border-grayC mb-4" @focus="blurAll()" />

						<div class="site__input-btn mt-10 flex flex-col md:flex-row items-center justify-center md:justify-end">
							<div class="w-full md:w-1/5">
								<NuxtLink to="/">
									<Button show-picture="nothing" name="Cancel"
										my-style="background: white; color: #000000; padding: 15px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2);" />
								</NuxtLink>
							</div>
							<button class="w-full md:w-1/5 ml-0 pt-6 md:pt-0 md:ml-6" type="submit">
								<Button
									my-style="background: #235365; font-family: 'Inter', sans-serif;padding: 15px;box-shadow: 0 1px 2px rgb(0 0 0 / 0.2);"
									name="Add" />
							</button>
						</div>
					</div>
				</form>
			</div>
		</main>
	</div>
</template>  