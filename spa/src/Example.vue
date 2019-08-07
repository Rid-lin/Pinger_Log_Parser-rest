<template>
  <div>
    <h1>{{message}}</h1>
    <app-ipnut :msg="message" @messChange="message = $event"></app-ipnut>
  </div>
</template>

<script>
import Input from './Field.vue';

export default {
    data() {
        return{
            message: 'Пробный текст'
        }
    },
  components: {
      'app-ipnut': Input
  }
}
</script>