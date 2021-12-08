package ai.verta.modeldb.experimentRun.subtypes;

import ai.verta.modeldb.common.exceptions.InternalErrorException;
import ai.verta.modeldb.common.futures.FutureJdbi;
import ai.verta.modeldb.common.subtypes.KeyValueHandler;
import java.util.concurrent.Executor;

public class AttributeHandler extends KeyValueHandler {
  public AttributeHandler(Executor executor, FutureJdbi jdbi, String entityName) {
    super(executor, jdbi, "attributes", entityName);
  }

  @Override
  protected String getTableName() {
    return "attribute";
  }

  @Override
  protected void setEntityIdReferenceColumn(String entityName) {
    switch (entityName) {
      case "ProjectEntity":
        this.entityIdReferenceColumn = "project_id";
        break;
      case "ExperimentRunEntity":
        this.entityIdReferenceColumn = "experiment_run_id";
        break;
      default:
        throw new InternalErrorException("Invalid entity name: " + entityName);
    }
  }
}
