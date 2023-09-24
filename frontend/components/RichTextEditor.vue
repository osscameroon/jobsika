<template>
  <client-only>
    <vue-editor :value="content" :editor-toolbar="toolbar" @input="$emit('input', $event)"/>
  </client-only>
</template>

<script>

export default {
  name: "RichTextEditor",
  props: {
    content: {
      type: String,
      default: ''
    },
    toolbar: {
      type: Array,
      default: () => []
    }
  }
}
</script>

<style lang="css">
  .ql-toolbar.ql-snow {
    background-color: #fff;
    border-radius: 0.375rem 0.375rem 0 0 !important;
  }

  .ql-container.ql-snow {
    border-radius: 0 0 0.375rem 0.375rem !important;
  }

  .ql-editor {
    background-color: #fff;
  }

  .ql-container.ql-snow:focus {
    background-color: #fff;
  }
</style>

