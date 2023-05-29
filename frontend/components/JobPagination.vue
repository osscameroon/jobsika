<template>
	<div class="bg-primary py-4 md:py-8 flex items-center justify-between sm:px-6">
		<div class="flex-1 flex justify-center sm:hidden">
			<span @click="startpage()" style="border-color: #919191" class="
				inline-flex
				relative
				px-3
				items-center
				rounded-md
				py-2
				bg-white
				border
				md:text-sm
				text-xs
				text-gray-500
				font-medium
				cursor-pointer
				hover:bg-gray-50
				mx-2
	    "
			>
				<span style="color: #919191" class="font-bold">&laquo;</span>
			</span>

			<span style="font-family: 'Inter', sans-serif" :class="`
				    relative
	          inline-flex
	          items-center
	          px-4
	          py-1
	          md:py-2
            border border-gray-300
	          text-xs
	          md:text-sm
	          font-medium
	          rounded-md
	          text-gray-700
	          cursor-pointer
            ${page === 1
					? 'bg-blue text-white border-none'
					: 'bg-blueDark text-white border-none'
				}
	            `" @click="previewpage()">
				{{$t("global_label_previous")}}
			</span>
			<span style="font-family: 'Inter', sans-serif" :class="`
						ml-3
						relative
						inline-flex
						items-center
						px-4
						py-1
						md:py-2
						text-xs
						md:text-sm
						font-medium
						rounded-md
						cursor-pointer
						${page === numberPage
					? 'bg-blue text-white border-none'
					: 'bg-blueDark text-white border-none'
				}
	        `" @click="nextpage()">
				{{$t("global_label_next")}}
			</span>
			<span @click="limitpage()" style="border-color: #919191" class="
	            relative
	            inline-flex
	            items-center
	            px-3
	            py-2
	            rounded-md
	            border
	            bg-white
	            text-xs
	            md:text-sm
	            font-medium
	            text-gray-500
	            hover:bg-gray-50
	            cursor-pointer
	            mx-2
	          ">
				<span style="color: #919191" class="font-bold">&raquo;</span>
			</span>
		</div>
		<div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
			<div>
				<p v-if="nbOfItems == 0" style="font-family: 'Inter', sans-serif" class="text-xs md:text-sm text-gray-700">
          {{$t("global_label_no_results")}}
				</p>
				<p v-else style="font-family: 'Inter', sans-serif" class="text-xs md:text-sm text-gray-700">
          {{$t("global_label_showing")}}
					<span class="font-medium">{{ (page - 1) * limit + 1 }}</span>
          {{$t("global_label_to")}}
					<span class="font-medium">{{ limit * (page - 1) + nbOfItems }}</span>
          {{$t("global_label_of")}}
					<span class="font-medium">{{ nbHits }}</span>
          {{$t("global_label_results")}}
				</p>
			</div>
			<div>
				<nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
					<span @click="startpage()" class="
	          relative
	          inline-flex
	          items-center
	          px-4
	          py-2
	          rounded-l-md
	          border border-gray-300
	          bg-white
	          text-xs
	          md:text-sm
	          font-medium
	          text-gray-500
	          hover:bg-gray-50
	          cursor-pointer
	        "
					>
						<span class="font-bold">&laquo;</span>
					</span>
					<div @click="previewpage()">
						<span :class="`z-10
	              bg-indigo-50
	              border-indigo-500
	              text-indigo-600
	              relative
	              inline-flex
	              items-center
	              px-2
	              py-2
	              border
							  text-xs
							  md:text-sm
							  font-medium
							  cursor-pointer
							  ${page === current1 ? 'bg-blue' : 'bg-white'}
							`"
						>
							<span class="sr-only">{{$t("global_label_previous")}}</span>
							<svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor"
								aria-hidden="true">
								<path fill-rule="evenodd"
									d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
									clip-rule="evenodd" />
							</svg>
						</span>
					</div>
					<div v-for="current1 in leftSide" :key="current1" @click="changepage(current1)">
						<span :class="`z-10
	            bg-indigo-50
	            border-indigo-500
	            text-indigo-600
	            relative
	            inline-flex
						  items-center
						  px-4
						  py-2
							border
						  text-xs
						  md:text-sm
						  font-medium
						  cursor-pointer
						  ${page === current1 ? 'bg-blue' : 'bg-white'}
					  `"
						>
							{{ current1 }}
						</span>
					</div>
					<div>
						<span class="
	          z-10
	          bg-blue
	          border-indigo-500
	          text-indigo-600
	          relative
	          inline-flex
	          items-center
	          px-4
	          py-2
	          border
	          text-xs
	          md:text-sm
	          font-medium
	          cursor-pointer
	        "
					>
							{{ page }}
						</span>
					</div>
					<div v-for="current in rightSide" :key="current" @click="changepage(current)">
						<span :class="`
							z-10
							bg-indigo-50
							border-indigo-500
							text-indigo-600
							relative
							inline-flex
							items-center
							px-4
							py-2
							border
							text-xs
							md:text-sm
							font-medium
							cursor-pointer
							${page === current1 ? 'bg-blue' : 'bg-white'}
						`"
						>
							{{ current }}
						</span>
					</div>
					<div @click="nextpage()">
						<span :class="`
							z-10
							bg-indigo-50
							border-indigo-500
							text-indigo-600
							relative
							inline-flex
							items-center
							px-2
							py-2
							border
							text-xs
							md:text-sm
							font-medium
							cursor-pointer
							${page === current1 ? 'bg-blue' : 'bg-white'}
						`"
						>
							<span class="sr-only">{{$t("global_label_next")}}</span>
							<svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor"
								aria-hidden="true">
								<path fill-rule="evenodd"
									d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
									clip-rule="evenodd" />
							</svg>
						</span>
					</div>
					<span @click="limitpage()" class="
	          relative
	          inline-flex
	          items-center
	          px-4
	          py-2
	          rounded-r-md
	          border border-gray-300
	          bg-white
	          text-xs
	          md:text-sm
	          font-medium
	          text-gray-500
	          hover:bg-gray-50
	          cursor-pointer
	        "
					>
						<span class="font-bold">&raquo;</span>
					</span>
				</nav>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	name: 'Jobs_Pagination_Number',
	data() {
		return {
			current1: '',
		}
	},
	computed: {
		page() {
			return this.$store.state.jobs.page
		},
		limit() {
			return this.$store.state.jobs.limit
		},
		nbHits() {
			return this.$store.state.jobs.nbHits
		},
		nbOfItems() {
			return this.$store.state.jobs.jobs.length
		},
		numberPage() {
			return Math.ceil(this.nbHits / this.limit)
		},
		leftSide() {
			const result = []
			for (let i = 5; i >= 1; i--) {
				if (this.page - i > 0) {
					result.push(this.page - i)
				}
			}
			return result
		},
		rightSide() {
			const result = []
			for (let i = 1; i <= 5; i++) {
				if (this.page + i < this.numberPage) {
					result.push(this.page + i)
				}
			}
			return result
		},
		filterjob() {
			return this.$store.state.ratings.filterjob
		},
		filtercompany() {
			return this.$store.state.ratings.filtercompany
		},
	},
	async created() {
		await this.fetchJobs();
	},
	methods: {
		async fetchJobs() {
			await this.$store.dispatch('getJobs', {
				page: this.page,
				limit: this.limit
			})
		},
		async setpage(value) {
			await this.$store.commit('jobs/SETPAGE', value)
		},
		changepage(value) {
			this.setpage(value)
			this.fetchJobs()
		},
		nextpage() {
			if (this.page + 1 <= this.numberPage) {
				this.changepage(this.page + 1)
			}
		},
		previewpage() {
			if (this.page - 1 > 0) {
				this.changepage(this.page - 1)
			}
		},
		startpage() {
			if (this.page === 1) {
				return null
			} else {
				this.changepage(1)
			}
		},
		limitpage() {
			if (this.numberPage === 1) {
				return null
			} else {
				this.changepage(this.numberPage)
			}
		},
	},
}
</script>
