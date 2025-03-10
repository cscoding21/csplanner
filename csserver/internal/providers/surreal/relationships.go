package surreal

/*

type Relationship struct {
	ID        string
	In        string
	Out       string
	CreatedAt RelationshipTime
}
type RelationshipTime struct {
	Written time.Time
}

func getRelationshipSQL(relationship string) string {
	return fmt.Sprintf("RELATE $node->%s->$edge SET time.written = time::now();", fmt.Sprint(relationship))
}

// RelateTo creates a relationship between two database objects
func (db *DBClient) RelateTo(ctx context.Context,
	nodeID string,
	edgeID string,
	relationship string) (*[]Relationship, error) {

	criteria := map[string]interface{}{
		"node": nodeID,
		"edge": edgeID,
	}

	sql := getRelationshipSQL(relationship)
	objectData, err := db.Client.Query(sql, criteria)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var outlist []Relationship
	_, err = surrealdb.UnmarshalRaw(objectData, &outlist)
	if err != nil {
		log.Error(err)
		return &outlist, err
	}

	return &outlist, nil
}

// TestRelationshipExist check to see if a relationship exists
func (db *DBClient) TestRelationshipExist(ctx context.Context, nodeID string, edgeID string, relationship string) (*Relationship, error) {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE in = $in and out = $out", string(relationship))
	criteria := map[string]interface{}{
		"in":  nodeID,
		"out": edgeID,
	}
	relationshipData, err := db.Client.Query(sql, criteria)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	outlist := make([]Relationship, 1)
	_, err = surrealdb.UnmarshalRaw(relationshipData, &outlist)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if len(outlist) == 0 || outlist[0].ID == "" {
		return nil, nil
	}

	return &outlist[0], nil
}

// GetRelationship gets a relationship by ID
func (db *DBClient) GetRelationship(ctx context.Context, relationshipID string) (*Relationship, error) {
	relationship, err := db.GetObjectById(relationshipID)
	if err != nil {
		return nil, err
	}

	output, err := marshal.SurrealUnmarshal[Relationship](relationship)
	if err != nil {
		return nil, err
	}

	return output, err
}

*/
