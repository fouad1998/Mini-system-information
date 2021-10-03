import { useContext, useCallback, useState } from "react";
import { ProviderContext } from "../provider/Provider";
import { Button, Col, Table, Row, Modal, Alert } from "antd";
import { DeleteFilled, DownloadOutlined, EditFilled } from "@ant-design/icons";
import EditEntry from "./EditEntry";

const { confirm } = Modal;
const Entries = () => {
  const { entries, loading, onRemoveHandler } = useContext(ProviderContext);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [alert, setAlert] = useState({ type: "error", message: "" });
  const [updateItemID, setUpdateItemID] = useState("");

  const onModalToggleHanlder = useCallback(() => {
    setIsModalVisible((state) => !state);
  }, []);

  const onRemoveEntryHandler = useCallback(
    (id, name) => {
      return () => {
        confirm({
          title: `Supprimer ${name}?`,
          content: `Etes vous sur de vouloir supprimer le fichier suivant: ${name}.`,
          okText: "Oui, je confirme",
          cancelText: "Annuler",
          onOk: async () => {
            const completed = await onRemoveHandler(id);
            if (completed) {
              setAlert({ type: "success", message: "Le fichier est supprimé avec succès" });
            } else {
              setAlert({ type: "error", message: "Fichier n'est pas supprimé" });
            }
          },
        });
      };
    },
    [onRemoveHandler]
  );

  const onEditEntryHandler = useCallback(
    (id, r) => {
      return () => {
        setUpdateItemID(id);
        onModalToggleHanlder();
      };
    },
    [onModalToggleHanlder]
  );

  return (
    <>
      <EditEntry
        visible={isModalVisible}
        id={updateItemID}
        setAlert={setAlert}
        onModalToggleHanlder={onModalToggleHanlder}
      />
      <Row gutter={[8, 16]}>
        {alert.message !== "" && (
          <Col span={24}>
            <Alert
              message={alert.message}
              type={alert.type}
              onClose={() => {
                setAlert((state) => ({ ...state, message: "" }));
              }}
              closable
            />
          </Col>
        )}
        <Col span={24}>
          <Table
            dataSource={entries}
            loading={loading}
            rowKey={(record) => record.id}
            title={() => <strong>Tous vous documents sont listés ci-dessous</strong>}
            pagination={{
              defaultPageSize: 25,
              disabled: false,
              position: ["bottomRight"],
              showTitle: true,
              pageSizeOptions: [25, 50, 100, 250, 500],
              showTotal: (total) => <strong>Nombre du documents enregistré est: {total}</strong>,
            }}
            columns={[
              {
                title: "Nom",
                onFilter(value, record) {
                  return record.name.indexOf(value) !== -1;
                },
                sorter(a, b) {
                  return a.name > b.name ? 1 : -1;
                },
                render(_, record) {
                  return record.name;
                },
              },
              {
                title: "Description",
                filtered: true,
                filterSearch: true,
                onFilter(value, record) {
                  return record.description.indexOf(value) !== -1;
                },
                render(_, record) {
                  return record.description;
                },
              },
              {
                title: "Proprietaire",
                onFilter(value, record) {
                  return record.owner.indexOf(value) !== -1;
                },
                render(_, record) {
                  return record.owner;
                },
              },
              {
                title: "Crée le",
                sorter(a, b) {
                  return a.createdAt > b.createdAt ? 1 : -1;
                },
                render(_, record) {
                  return new Date(record.createdAt).toLocaleDateString("fr-fr");
                },
              },
              {
                title: "Modifié le",
                sorter(a, b) {
                  return a.modifiedAt > b.modifiedAt ? 1 : -1;
                },
                render(_, record) {
                  return new Date(record.modifiedAt).toLocaleDateString("fr-fr");
                },
              },
              {
                title: "Extension",
                onFilter(value, record) {
                  return record.extension.indexOf(value) !== -1;
                },
                render(_, record) {
                  return record.extension;
                },
              },
              {
                title: "Nom fichier sur le system",
                onFilter(value, record) {
                  return record.filename.indexOf(value) !== -1;
                },
                render(_, record) {
                  return record.filename;
                },
              },
              {
                title: "Actions",
                render(_, record) {
                  return (
                    <Row gutter={[8, 8]} wrap={false}>
                      <Col>
                        <Button
                          shape="circle"
                          onClick={() => {
                            window.open("http://localhost:8000/file?id=" + record.id);
                          }}
                        >
                          <DownloadOutlined />
                        </Button>
                      </Col>
                      <Col>
                        <Button
                          onClick={onEditEntryHandler(record.id, record.name, record.description, record.owner)}
                          shape="circle"
                        >
                          <EditFilled />
                        </Button>
                      </Col>
                      <Col>
                        <Button shape="circle" onClick={onRemoveEntryHandler(record.id, record.name)} danger>
                          <DeleteFilled />
                        </Button>
                      </Col>
                    </Row>
                  );
                },
              },
            ]}
          />
        </Col>
      </Row>
    </>
  );
};

export default Entries;
