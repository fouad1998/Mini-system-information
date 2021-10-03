import { useContext, useCallback, useState, useEffect } from "react";
import { ProviderContext } from "../provider/Provider";
import { Col, Row, Modal, Input, Upload, AutoComplete } from "antd";
import { InboxOutlined } from "@ant-design/icons";

const { Dragger } = Upload;
const EditEntry = ({ visible, id, setAlert, onModalToggleHanlder }) => {
  const { entries, authors, onUpdateHandler } = useContext(ProviderContext);
  const [loadingUpdate, setLoadingUpdate] = useState(false);
  const [fileAdd, setFileAdd] = useState({
    name: "",
    description: "",
    owner: "",
    file: void 0,
  });

  const onResetHandler = useCallback(() => {
    setFileAdd({
      name: "",
      description: "",
      owner: "",
      file: void 0,
    });
  }, []);

  const onUpdateClikHandler = useCallback(async () => {
    setLoadingUpdate(true);
    const completed = await onUpdateHandler({ ...fileAdd, id });
    setLoadingUpdate(false);
    if (completed) {
      setAlert({ type: "success", message: "Fichier est mis à jour avec succès" });
      onResetHandler();
      onModalToggleHanlder();
    } else {
      setAlert({ type: "error", message: "Opération échaouie, le fichier n'est pas mis à jour" });
    }
  }, [fileAdd, id, onModalToggleHanlder, onResetHandler, onUpdateHandler, setAlert]);

  const onSetFileAddHandler = useCallback((key) => {
    return (event) => {
      const {
        target: { value },
      } = event;
      setFileAdd((state) => ({ ...state, [key]: value }));
    };
  }, []);

  useEffect(() => {
    const found = entries.find((e) => e.id === id);
    if (typeof found !== "object") {
      return;
    }

    console.log("foudn", found, found.owner);
    setFileAdd({ ...found });
  }, [entries, id]);

  console.log(fileAdd.owner);

  return (
    <Modal
      title="Ajouter un nouveau fichier"
      visible={visible}
      okButtonProps={{ disabled: loadingUpdate }}
      cancelButtonProps={{ disabled: loadingUpdate }}
      okText="Modifier"
      cancelText="Annuler"
      onCancel={onModalToggleHanlder}
      onOk={onUpdateClikHandler}
    >
      <Row gutter={[8, 8]}>
        <Col span={24}>
          <Input addonBefore={"Nom du fichier"} onChange={onSetFileAddHandler("name")} value={fileAdd.name} />
        </Col>
        <Col span={24}>
          <AutoComplete
            style={{ width: "100%" }}
            options={authors.map((author) => ({ label: author, value: author }))}
            value={fileAdd.owner}
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
  );
};

export default EditEntry;
