<template>
    <div>
        <ul class='tabs__header items-center justify-center'>
            <li v-for='(tab, index) in tabs' :key='tab.title' @click='selectTab(index)'
                :class='{ "tab__selected": (index == selectedIndex) }'>
                {{ tab.title }}
            </li>
        </ul>
        <slot></slot>
    </div>
</template>

<script>
export default {
    name: "TabsComponent",
    data() {
        return {
            selectedIndex: 0,
            tabs: []
        }
    },
    created() {
        this.tabs = this.$children
    },
    mounted() {
        this.selectTab(0)
    },
    methods: {
        selectTab(i) {
            this.selectedIndex = i
            this.tabs.forEach((tab, index) => {
                tab.isActive = (index === i)
            })
        }
    }
}
</script>

<style lang="css">
ul.tabs__header {
    /* display: block; */
    list-style: none;
    margin: 0 0 0 20px;
    padding: 0;
    display: flex;
}

ul.tabs__header>li {
    background-color: #e9ecf0;
    margin: 0;
    display: inline-block;
    margin-right: 32px;
    height: 52px;
    width: 268px;
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
}

ul.tabs__header>li.tab__selected {
    background-color: #d9d9d9;
    opacity: 1;
    font-weight: bold;
    color: #000;
}

.tab {
    display: inline-block;
    color: #235365;
    padding: 20px;
    min-width: 800px;
    border-radius: 10px;
    min-height: 400px;
}
</style>
