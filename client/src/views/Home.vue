<template>
  <v-container fluid>
    <v-slide-y-transition mode="out-in">
      <v-layout row wrap>
        <v-flex md3 sm6>
          <h1>{{bottomNav}}</h1>
          <v-card v-for="card in cards" flex>
            <v-card-title primary-title>
              <h3 class="headline mb-0" v-text="card.title"></h3></v-card-title>
            <v-card-text v-text="card.text"></v-card-text>
            <v-card-actions>
              <v-btn icon><v-icon color="green" @click="upvote = true">thumb_up</v-icon></v-btn>
            </v-card-actions>
          </v-card>
          <v-snackbar v-model="upvote">
            Sikeres upvote
            <v-btn dark flat @click="upvote = false">
              Close
            </v-btn>
          </v-snackbar>
        </v-flex>
        <v-bottom-nav :active.sync="bottomNav" :value="true" absolute color="white">
          <v-btn color="primary" flat value="Témák">
            <span>Témák</span>
            <v-icon>history</v-icon>
          </v-btn>

          <v-btn color="primary" flat value="Sprintek">
            <span>Sprintek</span>
            <v-icon>favorite</v-icon>
          </v-btn>

          <v-btn color="primary" flat value="Saját témáim">
            <span>Saját témáim</span>
            <v-icon>place</v-icon>
          </v-btn>
        </v-bottom-nav>
      </v-layout>
    </v-slide-y-transition>
  </v-container>
</template>

<script>
  export default {
    data () {
      return {
        upvote: false,
        cards: [
          {
            title: 'Első retro téma',
            text: 'Jobb az app, mint a csatosüveg'
          },
          {
            title: 'Második téma',
            text: 'Majd jobb lesz'
          }
        ],
        bottomNav: 'Témák'
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
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
.v-card {
  margin-bottom: 10px;
}
</style>
