/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CommentType } from './CommentType';
import {
    CommentTypeFromJSON,
    CommentTypeFromJSONTyped,
    CommentTypeToJSON,
} from './CommentType';

/**
 * Comment related to the profile/reservation.
 * @export
 * @interface StagedProfileCommentInfoType
 */
export interface StagedProfileCommentInfoType {
    /**
     * 
     * @type {CommentType}
     * @memberof StagedProfileCommentInfoType
     */
    comment?: CommentType;
    /**
     * Time stamp of the creation.
     * @type {Date}
     * @memberof StagedProfileCommentInfoType
     */
    createDateTime?: Date;
    /**
     * ID of creator. The creator could be a software system identifier or an identifier of an employee resposible for the creation.
     * @type {string}
     * @memberof StagedProfileCommentInfoType
     */
    creatorId?: string;
    /**
     * Time stamp of last modification.
     * @type {Date}
     * @memberof StagedProfileCommentInfoType
     */
    lastModifyDateTime?: Date;
    /**
     * Identifies the last software system or person to modify a record.
     * @type {string}
     * @memberof StagedProfileCommentInfoType
     */
    lastModifierId?: string;
    /**
     * The line number of the comment.
     * @type {number}
     * @memberof StagedProfileCommentInfoType
     */
    lineNo?: number;
    /**
     * Indicates comment inactive date.
     * @type {Date}
     * @memberof StagedProfileCommentInfoType
     */
    inActiveDate?: Date;
    /**
     * The error in a user defined field in a staged profile with an invalid status.
     * @type {string}
     * @memberof StagedProfileCommentInfoType
     */
    errorDescription?: string;
    /**
     * Indicates whether the comment information is valid.
     * @type {boolean}
     * @memberof StagedProfileCommentInfoType
     */
    valid?: boolean;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof StagedProfileCommentInfoType
     */
    allRowsFetched?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof StagedProfileCommentInfoType
     */
    totalRows?: number;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof StagedProfileCommentInfoType
     */
    id?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof StagedProfileCommentInfoType
     */
    type?: string;
}

/**
 * Check if a given object implements the StagedProfileCommentInfoType interface.
 */
export function instanceOfStagedProfileCommentInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function StagedProfileCommentInfoTypeFromJSON(json: any): StagedProfileCommentInfoType {
    return StagedProfileCommentInfoTypeFromJSONTyped(json, false);
}

export function StagedProfileCommentInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): StagedProfileCommentInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'comment': !exists(json, 'comment') ? undefined : CommentTypeFromJSON(json['comment']),
        'createDateTime': !exists(json, 'createDateTime') ? undefined : (new Date(json['createDateTime'])),
        'creatorId': !exists(json, 'creatorId') ? undefined : json['creatorId'],
        'lastModifyDateTime': !exists(json, 'lastModifyDateTime') ? undefined : (new Date(json['lastModifyDateTime'])),
        'lastModifierId': !exists(json, 'lastModifierId') ? undefined : json['lastModifierId'],
        'lineNo': !exists(json, 'lineNo') ? undefined : json['lineNo'],
        'inActiveDate': !exists(json, 'inActiveDate') ? undefined : (new Date(json['inActiveDate'])),
        'errorDescription': !exists(json, 'errorDescription') ? undefined : json['errorDescription'],
        'valid': !exists(json, 'valid') ? undefined : json['valid'],
        'allRowsFetched': !exists(json, 'allRowsFetched') ? undefined : json['allRowsFetched'],
        'totalRows': !exists(json, 'totalRows') ? undefined : json['totalRows'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'type': !exists(json, 'type') ? undefined : json['type'],
    };
}

export function StagedProfileCommentInfoTypeToJSON(value?: StagedProfileCommentInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'comment': CommentTypeToJSON(value.comment),
        'createDateTime': value.createDateTime === undefined ? undefined : (value.createDateTime.toISOString()),
        'creatorId': value.creatorId,
        'lastModifyDateTime': value.lastModifyDateTime === undefined ? undefined : (value.lastModifyDateTime.toISOString()),
        'lastModifierId': value.lastModifierId,
        'lineNo': value.lineNo,
        'inActiveDate': value.inActiveDate === undefined ? undefined : (value.inActiveDate.toISOString().substr(0,10)),
        'errorDescription': value.errorDescription,
        'valid': value.valid,
        'allRowsFetched': value.allRowsFetched,
        'totalRows': value.totalRows,
        'id': value.id,
        'type': value.type,
    };
}

