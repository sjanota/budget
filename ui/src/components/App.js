import React from 'react';
import {useQuery} from "@apollo/react-hooks";
import {gql} from 'apollo-boost';


const GET_EXPENSES = gql`
    query {
        expenses {
            id
            title
            date
            total
            location
            account {
                id
                name
                available
            }
            entries {
                category {
                    expenses(since: "123") {
                        total
                    }
                }
            }
        }
    }`;

function App() {
  const {loading, error, data} = useQuery(GET_EXPENSES);

  if (loading) return <p>Loading...</p>;
  if (error) {
    console.warn(error);
    return <p>Error :(</p>;
  }

  console.log(loading, error, data);

  return (
    <div className="App">
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
}

export default App;
