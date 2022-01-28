export { default as Average } from '../../components/Average.vue'
export { default as Button } from '../../components/Button.vue'
export { default as Company } from '../../components/Company.vue'
export { default as Header } from '../../components/Header.vue'
export { default as MyInput } from '../../components/MyInput.vue'
export { default as Pagination } from '../../components/Pagination.vue'
export { default as Select } from '../../components/Select.vue'
export { default as Title } from '../../components/Title.vue'

// nuxt/nuxt.js#8607
function wrapFunctional(options) {
  if (!options || !options.functional) {
    return options
  }

  const propKeys = Array.isArray(options.props) ? options.props : Object.keys(options.props || {})

  return {
    render(h) {
      const attrs = {}
      const props = {}

      for (const key in this.$attrs) {
        if (propKeys.includes(key)) {
          props[key] = this.$attrs[key]
        } else {
          attrs[key] = this.$attrs[key]
        }
      }

      return h(options, {
        on: this.$listeners,
        attrs,
        props,
        scopedSlots: this.$scopedSlots,
      }, this.$slots.default)
    }
  }
}
