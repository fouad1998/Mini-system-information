import { Button, Col, Row, Modal, Input, AutoComplete, Upload, Alert } from "antd";
import { FileAddFilled, InboxOutlined } from "@ant-design/icons";
import { useCallback, useContext, useState } from "react";
import { ProviderContext } from "../provider/Provider";

const { Dragger } = Upload;

const AddEntry = () => {
  const { authors, onAddHandler, onSearch } = useContext(ProviderContext);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [alert, setAlert] = useState({ type: "error", message: "" });
  const [fileAdd, setFileAdd] = useState({
    name: "",
    description: "",
    owner: "",
    file: void 0,
  });
  const [loading, setLoading] = useState(false);
  const onSetFileAddHandler = useCallback((key) => {
    return (event) => {
      const {
        target: { value },
      } = event;
      setFileAdd((state) => ({ ...state, [key]: value }));
    };
  }, []);

  const onModalToggleHanlder = useCallback(() => {
    setIsModalVisible((state) => !state);
  }, []);

  const onResetHandler = useCallback(() => {
    setFileAdd({
      name: "",
      description: "",
      owner: "",
      file: void 0,
    });
  }, []);

  const onAddClikHandler = useCallback(async () => {
    setLoading(true);
    const completed = await onAddHandler(fileAdd);
    setLoading(false);
    if (completed) {
      setAlert({ type: "success", message: "Fichier est ajouté avec succès" });
      onResetHandler();
      onModalToggleHanlder();
    } else {
      setAlert({ type: "error", message: "Opération échaouie, le fichier n'est pas ajouté" });
    }
  }, [fileAdd, onAddHandler, onModalToggleHanlder, onResetHandler]);

  return (
    <Row gutter={[8, 8]} justify="end">
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
      <Col flex="1 1 auto">
        <Input.Search size="large" placeholder="chercher parmi vos documents et proprietaires" onSearch={onSearch} />
      </Col>
      <Col flex="1 1 auto"></Col>
      <Col>
        <Button type="primary" size="large" onClick={onModalToggleHanlder}>
          <Row gutter={8} wrap={false}>
            <Col>
              <FileAddFilled />
            </Col>
            <Col>Add</Col>
          </Row>
        </Button>
      </Col>
      <Modal
        title="Ajouter un nouveau fichier"
        visible={isModalVisible}
        okButtonProps={{ disabled: loading }}
        cancelButtonProps={{ disabled: loading }}
        okText="Ajouter"
        cancelText="Annuler"
        onCancel={onModalToggleHanlder}
        onOk={onAddClikHandler}
      >
        <Row gutter={[8, 8]}>
          <Col span={24}>
            <Input addonBefore={"Nom du fichier"} onChange={onSetFileAddHandler("name")} value={fileAdd.name} />
          </Col>
          <Col span={24}>
            <AutoComplete
              style={{ width: "100%" }}
              value={fileAdd.owner}
              options={authors.map((author) => ({ label: author, value: author }))}
              onSelect={(value) => {
                setFileAdd((state) => ({ ...state, owner: value }));
              }}
            >
              <Input addonBefore={"Proprietaire"} onChange={onSetFileAddHandler("owner")} value={fileAdd.owner} />
            </AutoComplete>
          </Col>
          <Col span={24}>
            <Input.TextArea
              rows={4}
              autoSize={{ maxRows: 10, minRows: 4 }}
              placeholder={"Description"}
              onChange={onSetFileAddHandler("description")}
              value={fileAdd.description}
            />
          </Col>
          <Col span={24}>
            <Dragger
              name="file"
              fileList={typeof fileAdd.file === "undefined" ? [] : [fileAdd.file]}
              onRemove={() => {
                setFileAdd((state) => ({ ...state, file: void 0 }));
                return true;
              }}
              beforeUpload={(file) => {
                setFileAdd((state) => ({ ...state, file }));
                return false;
              }}
            >
              <p className="ant-upload-drag-icon">
                <InboxOutlined />
              </p>
              <p className="ant-upload-text">Cliquez ou lachez un fichier dans la zone</p>
            </Dragger>
          </Col>
        </Row>
      </Modal>
    </Row>
  );
};

export default AddEntry;
