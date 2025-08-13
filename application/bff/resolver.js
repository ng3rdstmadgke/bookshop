export const resolvers = {
  Query: {
    book: async (parent, args, context) => {
      // const response = await context.dataSources.catalogueApi.getBook(args.id)
      const response = books[args.id - 1]
      return response
    },
    books: async (parent, args, context) => {
      // const response = await context.dataSources.catalogueApi.listBooks()
      const response = books
      return response
    }
  }
}

const books = [
  {id: 1, title: 'The Awakening', author: 'Kate Chopin', price: 1000},
  {id: 2, title: 'hogehoge', author: 'abcdefg', price: 2000}
]