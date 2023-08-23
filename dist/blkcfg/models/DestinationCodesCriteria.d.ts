/**
 * OPERA Cloud Block Configuration API
 * APIs for Block configuration, such as creating, updating, fetching and removing codes related to blocks. <br />< This might include fetching the block cancellation reasons, or creating new block refused reasons.  Wash schedules can be create, or new reservation methods could be added for a property.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { DestinationCodeType } from './DestinationCodeType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request object for creating Destination Codes.
 * @export
 * @interface DestinationCodesCriteria
 */
export interface DestinationCodesCriteria {
    /**
     * List of Destination Codes.
     * @type {Array<DestinationCodeType>}
     * @memberof DestinationCodesCriteria
     */
    destinationCodes?: Array<DestinationCodeType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof DestinationCodesCriteria
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof DestinationCodesCriteria
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the DestinationCodesCriteria interface.
 */
export declare function instanceOfDestinationCodesCriteria(value: object): boolean;
export declare function DestinationCodesCriteriaFromJSON(json: any): DestinationCodesCriteria;
export declare function DestinationCodesCriteriaFromJSONTyped(json: any, ignoreDiscriminator: boolean): DestinationCodesCriteria;
export declare function DestinationCodesCriteriaToJSON(value?: DestinationCodesCriteria | null): any;
