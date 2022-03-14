from seldon_core.user_model import SeldonComponent
import numpy as np
from typing import Dict, List, Union, Iterable
import logging
import torch
import os

logger = logging.getLogger(__name__)

class PytorchServer(SeldonComponent):
        def __init__(self, model_uri: str = None, model_volume_path: str = None, model_user_Id: str = None,
                     model_Id: str = None, model_version: str = None ):
            super().__init__()
            self.model_uri = model_uri
            self.model_volume_path = model_volume_path
            self.model_userId = model_user_Id
            self.model_Id = model_Id
            self.model_version = model_version
            self.ready = False
            self.model = None
            self.load()

        def load(self):
            logger.info("-----------2.0.1-----load model--------------")
            user_model_from_path = os.path.join(self.model_volume_path, self.model_userId, self.model_Id,
                                                self.model_version, "model.pth")
            logger.info(f"model full path in docker is: {user_model_from_path}....")
            model = torch.load(user_model_from_path)
            # model = model.double()
            self.model = model.eval()
            self.ready = True
            logger.info(f"model is loaded, waiting to predict")

        def predict(self, X: np.ndarray, names: Iterable[str], meta: Dict = None) -> Union[np.ndarray, List, str, bytes]:
            try:
                logger.info(f"------------predict method, try to do predict--------------")
                logger.info(f"1st: transform list to array")
                X = np.array(X,np.float32)
                logger.info(f"2nd: transform array to tensor")
                tensor = torch.from_numpy(X)
                logger.info(f"3rd: forward to predict")
                outputs = self.model.forward(tensor.type(torch.float32))
                logger.info(f"-------------predict successfully!----------------------")
                logger.info(f"4th: infer result is: {outputs.tolist()}")
                return outputs.tolist()
                return "test predict method"
            except Exception as ex:
                logging.exception("Exception during predict")
                return "model predicted failed, please check your request parameters!"