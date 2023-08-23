/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ChangeRelationsType } from './ChangeRelationsType';
import {
    ChangeRelationsTypeFromJSON,
    ChangeRelationsTypeFromJSONTyped,
    ChangeRelationsTypeToJSON,
} from './ChangeRelationsType';
import type { MasterAccountInfoType } from './MasterAccountInfoType';
import {
    MasterAccountInfoTypeFromJSON,
    MasterAccountInfoTypeFromJSONTyped,
    MasterAccountInfoTypeToJSON,
} from './MasterAccountInfoType';
import type { RelationshipProfileType } from './RelationshipProfileType';
import {
    RelationshipProfileTypeFromJSON,
    RelationshipProfileTypeFromJSONTyped,
    RelationshipProfileTypeToJSON,
} from './RelationshipProfileType';

/**
 * Relationship Type contains information about the associations between and among individuals, companies, travel agents, groups, sources, and contact profiles.
 * @export
 * @interface RelationshipInfoType
 */
export interface RelationshipInfoType {
    /**
     * 
     * @type {ChangeRelationsType}
     * @memberof RelationshipInfoType
     */
    changeRelationship?: ChangeRelationsType;
    /**
     * 
     * @type {RelationshipProfileType}
     * @memberof RelationshipInfoType
     */
    relationshipProfile?: RelationshipProfileType;
    /**
     * 
     * @type {MasterAccountInfoType}
     * @memberof RelationshipInfoType
     */
    masterAccountInfo?: MasterAccountInfoType;
    /**
     * Relationship identifier.
     * @type {string}
     * @memberof RelationshipInfoType
     */
    id?: string;
    /**
     * Indicates the type of relationship the current profile(Source Profile) has with the related profile(Target Profile).
     * @type {string}
     * @memberof RelationshipInfoType
     */
    relation?: string;
    /**
     * Displays the description of relationship the current profile(Source Profile) has with the related profile(Target Profile).
     * @type {string}
     * @memberof RelationshipInfoType
     */
    relationDescription?: string;
    /**
     * Displays the type of relationship the Related profile(Target Profile) has with the current profile(Source Profile).
     * @type {string}
     * @memberof RelationshipInfoType
     */
    targetRelation?: string;
    /**
     * Displays the description of the target relation(Target Profile).
     * @type {string}
     * @memberof RelationshipInfoType
     */
    targetRelationDescription?: string;
}

/**
 * Check if a given object implements the RelationshipInfoType interface.
 */
export function instanceOfRelationshipInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RelationshipInfoTypeFromJSON(json: any): RelationshipInfoType {
    return RelationshipInfoTypeFromJSONTyped(json, false);
}

export function RelationshipInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RelationshipInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'changeRelationship': !exists(json, 'changeRelationship') ? undefined : ChangeRelationsTypeFromJSON(json['changeRelationship']),
        'relationshipProfile': !exists(json, 'relationshipProfile') ? undefined : RelationshipProfileTypeFromJSON(json['relationshipProfile']),
        'masterAccountInfo': !exists(json, 'masterAccountInfo') ? undefined : MasterAccountInfoTypeFromJSON(json['masterAccountInfo']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'relation': !exists(json, 'relation') ? undefined : json['relation'],
        'relationDescription': !exists(json, 'relationDescription') ? undefined : json['relationDescription'],
        'targetRelation': !exists(json, 'targetRelation') ? undefined : json['targetRelation'],
        'targetRelationDescription': !exists(json, 'targetRelationDescription') ? undefined : json['targetRelationDescription'],
    };
}

export function RelationshipInfoTypeToJSON(value?: RelationshipInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'changeRelationship': ChangeRelationsTypeToJSON(value.changeRelationship),
        'relationshipProfile': RelationshipProfileTypeToJSON(value.relationshipProfile),
        'masterAccountInfo': MasterAccountInfoTypeToJSON(value.masterAccountInfo),
        'id': value.id,
        'relation': value.relation,
        'relationDescription': value.relationDescription,
        'targetRelation': value.targetRelation,
        'targetRelationDescription': value.targetRelationDescription,
    };
}

