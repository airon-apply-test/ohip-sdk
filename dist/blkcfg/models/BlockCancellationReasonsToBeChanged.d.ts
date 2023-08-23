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
import type { BlockCancellationReasonType } from './BlockCancellationReasonType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 * Request object for changing Block Cancellation Reasons.
 * @export
 * @interface BlockCancellationReasonsToBeChanged
 */
export interface BlockCancellationReasonsToBeChanged {
    /**
     * List of Block Cancellation Reasons.
     * @type {Array<BlockCancellationReasonType>}
     * @memberof BlockCancellationReasonsToBeChanged
     */
    blockCancellationReasons?: Array<BlockCancellationReasonType>;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof BlockCancellationReasonsToBeChanged
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof BlockCancellationReasonsToBeChanged
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the BlockCancellationReasonsToBeChanged interface.
 */
export declare function instanceOfBlockCancellationReasonsToBeChanged(value: object): boolean;
export declare function BlockCancellationReasonsToBeChangedFromJSON(json: any): BlockCancellationReasonsToBeChanged;
export declare function BlockCancellationReasonsToBeChangedFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockCancellationReasonsToBeChanged;
export declare function BlockCancellationReasonsToBeChangedToJSON(value?: BlockCancellationReasonsToBeChanged | null): any;
