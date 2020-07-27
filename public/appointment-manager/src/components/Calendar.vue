<template>
  <div class="hello">
    <div class="calendar">
    <vc-calendar    mode="range"
                    :value="null"
                    color="red"
                    is-dark
                    is-inline
                    :columns="layout.columns"
                    :rows="layout.rows"
                    :is-expanded="layout.isExpanded"
                    
    />
  </div>

  </div>
</template>

<script>
export default {
  computed: {
    layout() {
      return this.$screens( //this function adjust the layout to be responsive
        {
          // Default layout for mobile
          default: {
            columns: 1,
            rows: 1,
            isExpanded: true,
          },
          md: {
            columns: 1,
            rows: 1,
            isExpanded: true,
          },
          lg: {
            columns: 1,
            rows: 1,
            isExpanded: true,
          },
          // Override for large screens
          xl: {
            columns: 1,
            rows: 1,
            isExpanded: true,
          },
        },
      );
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
