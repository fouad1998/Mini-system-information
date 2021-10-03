import { Alert, Row, Col, Button } from "antd";
import axios from "axios";
import { createContext, useCallback, useEffect, useState } from "react";

export const ProviderContext = createContext({});
const Provider = ({ children }) => {
  const [loading, setLoading] = useState(false);
  const [search, setSearch] = useState("");
  const [entries, setEntries] = useState([]);
  const [authors, setAuthors] = useState([]);
  const [error, setError] = useState(void 0);

  const onSearch = useCallback((value) => {
    setSearch(value);
  }, []);

  const onFetchHandler = useCallback(async () => {
    try {
      setError(void 0);
      setLoading(true);
      const { data } = await axios.get("http://localhost:8000/entries");
      const { entries, authors } = data;
      setAuthors(authors);
      setEntries(entries);
    } catch (e) {
      console.error(e);
      setError("Échec de téléchargment des données, svp reessayer une autre fois");
    } finally {
      setLoading(false);
    }
  }, []);

  const onAddHandler = useCallback(async (add) => {
    try {
      const dataF = new FormData();
      dataF.append("name", add.name);
      dataF.append("description", add.description);
      dataF.append("owner", add.owner);
      dataF.append("file", add.file);
      const { data: entry } = await axios.post("http://localhost:8000/add", dataF);
      setEntries((state) => [...state, entry]);
      setAuthors((state) => {
        if (state.includes(entry.owner)) {
          return state;
        }

        return [...state, entry.owner];
      });
      return true;
    } catch {
      return false;
    }
  }, []);

  const onRemoveHandler = useCallback(async (id) => {
    try {
      await axios.delete("http://localhost:8000/remove?id=" + id);
      setEntries((state) => state.filter((entry) => entry.id !== id));
      return true;
    } catch (error) {
      console.error(error);
      return false;
    }
  }, []);

  const onUpdateHandler = useCallback(async (data) => {
    try {
      const f = new FormData();
      f.append("id", data.id);
      f.append("name", data.name);
      f.append("description", data.description);
      f.append("owner", data.owner);
      if (typeof data.file === "object") {
        f.append("file", data.file);
      }

      await axios.post("http://localhost:8000/update", f);
      setEntries((state) => {
        const found = state.find((e) => e.id === data.id);
        if (typeof found !== "object") {
          return state;
        }

        found.name = data.name;
        found.description = data.description;
        found.owner = data.owner;
        return [...state];
      });
      return true;
    } catch (error) {}
  }, []);

  useEffect(() => {
    onFetchHandler();
  }, [onFetchHandler]);

  return (
    <ProviderContext.Provider
      value={{
        loading,
        entries: entries.filter((entry) =>
          Object.values(entry)
            .map((value) => value.toString().indexOf(search) !== -1)
            .reduce((prev, curr) => prev || curr, false)
        ),
        authors,
        onAddHandler,
        onSearch,
        onUpdateHandler,
        onRemoveHandler,
      }}
    >
      <Row>
        {error && (
          <Col span={24}>
            <Alert message={error} type="error" action={<Button onClick={onFetchHandler}>Recharger</Button>} />
          </Col>
        )}
        <Col span={24}>{children}</Col>
      </Row>
    </ProviderContext.Provider>
  );
};

export default Provider;
