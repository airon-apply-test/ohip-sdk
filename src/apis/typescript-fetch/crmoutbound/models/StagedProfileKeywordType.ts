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
import type { KeywordDetailType } from './KeywordDetailType';
import {
    KeywordDetailTypeFromJSON,
    KeywordDetailTypeFromJSONTyped,
    KeywordDetailTypeToJSON,
} from './KeywordDetailType';

/**
 * 
 * @export
 * @interface StagedProfileKeywordType
 */
export interface StagedProfileKeywordType {
    /**
     * 
     * @type {KeywordDetailType}
     * @memberof StagedProfileKeywordType
     */
    keywordDetail?: KeywordDetailType;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof StagedProfileKeywordType
     */
    type?: string;
    /**
     * The keyword value.
     * @type {string}
     * @memberof StagedProfileKeywordType
     */
    keyword?: string;
    /**
     * The error in keyword information in a staged profile with an invalid status
     * @type {string}
     * @memberof StagedProfileKeywordType
     */
    errorDescription?: string;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof StagedProfileKeywordType
     */
    id?: string;
}

/**
 * Check if a given object implements the StagedProfileKeywordType interface.
 */
export function instanceOfStagedProfileKeywordType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function StagedProfileKeywordTypeFromJSON(json: any): StagedProfileKeywordType {
    return StagedProfileKeywordTypeFromJSONTyped(json, false);
}

export function StagedProfileKeywordTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): StagedProfileKeywordType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'keywordDetail': !exists(json, 'keywordDetail') ? undefined : KeywordDetailTypeFromJSON(json['keywordDetail']),
        'type': !exists(json, 'type') ? undefined : json['type'],
        'keyword': !exists(json, 'keyword') ? undefined : json['keyword'],
        'errorDescription': !exists(json, 'errorDescription') ? undefined : json['errorDescription'],
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function StagedProfileKeywordTypeToJSON(value?: StagedProfileKeywordType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'keywordDetail': KeywordDetailTypeToJSON(value.keywordDetail),
        'type': value.type,
        'keyword': value.keyword,
        'errorDescription': value.errorDescription,
        'id': value.id,
    };
}

