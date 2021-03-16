<template>
    <apexchart type="radialBar" :options="optiions" :series="series"></apexchart>
</template>


<script>
export default {
    data() {
        return {
            series: [0],
            options: {
                labels: ['CPU Usage']
            }
         };
      },
      mounted: function(){
          wails.Events.on("cpu_usage", cpu_usage => {
              if (cpu_usage) {
                  this.series = [cpu_usage.avg];
              }
          });
      }
    };
</script>
